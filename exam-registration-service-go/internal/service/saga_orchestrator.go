package service

import (
	"exam-registration-service/internal/domain"
	"exam-registration-service/internal/messaging"
	"exam-registration-service/internal/repository"
	"log"
	"time"
)

type SagaOrchestrator struct {
	regRepo  repository.ExamRegistrationRepository
	itemRepo repository.SagaItemRepository
	publisher *messaging.Publisher
}

func NewSagaOrchestrator(
	regRepo repository.ExamRegistrationRepository,
	itemRepo repository.SagaItemRepository,
	publisher *messaging.Publisher,
) *SagaOrchestrator {
	return &SagaOrchestrator{
		regRepo:   regRepo,
		itemRepo:  itemRepo,
		publisher: publisher,
	}
}

func (s *SagaOrchestrator) StartExamRegistration(studentID int, examID uint, courseID int) (*domain.ExamRegistration, error) {
	// Create exam registration
	registration := &domain.ExamRegistration{
		StudentID: studentID,
		ExamID:    examID,
		CourseID:  courseID,
		Status:    domain.StatusInitialized,
	}

	if err := s.regRepo.Create(registration); err != nil {
		return nil, err
	}

	// Create saga items
	for _, def := range domain.SagaItemDefinitions {
		item := &domain.SagaItem{
			Type:               def.Type,
			SagaType:           def.SagaType,
			ExecutionOrder:     def.ExecutionOrder,
			Status:             domain.StatusInitialized,
			ExamRegistrationID: registration.ID,
			ReturnedPayload:    make(domain.JSON),
		}

		if err := s.itemRepo.Create(item); err != nil {
			log.Printf("Failed to create saga item: %v", err)
			continue
		}
	}

	// Update registration status
	registration.Status = domain.StatusInProgress
	s.regRepo.Update(registration)

	// Start executing saga
	go s.executeSaga(registration.ID)

	return registration, nil
}

func (s *SagaOrchestrator) executeSaga(registrationID uint) {
	log.Printf("Starting saga execution for registration %d", registrationID)

	registration, err := s.regRepo.FindByID(registrationID)
	if err != nil {
		log.Printf("Failed to find registration: %v", err)
		return
	}

	items, err := s.itemRepo.FindByExamRegistrationID(registrationID)
	if err != nil {
		log.Printf("Failed to find saga items: %v", err)
		return
	}

	for _, item := range items {
		log.Printf("Executing saga item %d: %s", item.ID, item.SagaType)

		now := time.Now()
		item.Status = domain.StatusInProgress
		item.StartedAt = &now
		s.itemRepo.Update(&item)

		// Publish message to appropriate exchange
		if err := s.publishSagaMessage(&item, registration); err != nil {
			log.Printf("Failed to publish saga message: %v", err)
			s.handleSagaFailure(registration, &item)
			return
		}

		// Wait for response (in production, this would be event-driven)
		// For now, we'll just mark it as in progress and let the consumer handle responses
		time.Sleep(2 * time.Second)
	}
}

func (s *SagaOrchestrator) publishSagaMessage(item *domain.SagaItem, registration *domain.ExamRegistration) error {
	switch item.SagaType {
	case domain.SagaTypeUserClassVerification:
		return s.publisher.PublishUserClassVerification(item.ID, registration.StudentID, registration.CourseID)
	case domain.SagaTypeUserWalletValidation:
		return s.publisher.PublishUserWalletValidation(item.ID, registration.StudentID, registration.ExamID)
	case domain.SagaTypeUserWalletInsert:
		return s.publisher.PublishUserWalletInsert(item.ID, registration.StudentID, registration.ExamID)
	case domain.SagaTypeExamRegistration:
		return s.publisher.PublishExamRegistration(item.ID, registration.StudentID, registration.ExamID, registration.ID)
	default:
		log.Printf("Unknown saga type: %s", item.SagaType)
		return nil
	}
}

func (s *SagaOrchestrator) HandleSagaResponse(sagaItemID uint, success bool, message string) error {
	item, err := s.itemRepo.FindByID(sagaItemID)
	if err != nil {
		return err
	}

	now := time.Now()
	item.FinishedAt = &now
	item.ReturnedPayload = domain.JSON{
		"success": success,
		"message": message,
	}

	if success {
		item.Status = domain.StatusFinished
		s.itemRepo.Update(item)

		// Check if all items are finished
		items, _ := s.itemRepo.FindByExamRegistrationID(item.ExamRegistrationID)
		allFinished := true
		for _, i := range items {
			if i.Status != domain.StatusFinished {
				allFinished = false
				break
			}
		}

		if allFinished {
			registration, _ := s.regRepo.FindByID(item.ExamRegistrationID)
			registration.Status = domain.StatusFinished
			s.regRepo.Update(registration)
			log.Printf("Saga completed successfully for registration %d", item.ExamRegistrationID)
		}
	} else {
		item.Status = "failed"
		s.itemRepo.Update(item)

		// Handle failure
		registration, _ := s.regRepo.FindByID(item.ExamRegistrationID)
		s.handleSagaFailure(registration, item)
	}

	return nil
}

func (s *SagaOrchestrator) handleSagaFailure(registration *domain.ExamRegistration, failedItem *domain.SagaItem) {
	log.Printf("Saga failed at step %s for registration %d", failedItem.SagaType, registration.ID)
	registration.Status = "failed"
	s.regRepo.Update(registration)

	// Here you would implement compensating transactions to rollback changes
	// For now, we'll just log the failure
}
