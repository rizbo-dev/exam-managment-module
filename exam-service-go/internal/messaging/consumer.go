package messaging

import (
	"encoding/json"
	"exam-service/internal/domain"
	"exam-service/internal/repository"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn              *amqp.Connection
	channel           *amqp.Channel
	examRepo          repository.ExamRepository
	examStudentRepo   repository.ExamStudentRepository
	publisher         *Publisher
}

func NewConsumer(rabbitmqURL string, examRepo repository.ExamRepository, examStudentRepo repository.ExamStudentRepository, publisher *Publisher) (*Consumer, error) {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn:            conn,
		channel:         channel,
		examRepo:        examRepo,
		examStudentRepo: examStudentRepo,
		publisher:       publisher,
	}, nil
}

func (c *Consumer) StartConsuming() error {
	// Declare exchange
	err := c.channel.ExchangeDeclare(
		"exam-registration",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Declare queue
	queue, err := c.channel.QueueDeclare(
		"exam-registration",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Bind queue to exchange
	err = c.channel.QueueBind(
		queue.Name,
		"",
		"exam-registration",
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Start consuming
	msgs, err := c.channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			c.handleExamRegistrationMessage(msg)
		}
	}()

	log.Println("Started consuming exam registration messages")
	return nil
}

func (c *Consumer) handleExamRegistrationMessage(msg amqp.Delivery) {
	var message struct {
		ExamID            uint `json:"examId"`
		StudentID         int  `json:"studentId"`
		ExamRegistrationID uint `json:"examRegistrationId"`
		SagaItemID        uint `json:"sagaItemId"`
	}

	if err := json.Unmarshal(msg.Body, &message); err != nil {
		log.Printf("Failed to unmarshal message: %v", err)
		msg.Nack(false, false)
		return
	}

	log.Printf("Processing exam registration: ExamID=%d, StudentID=%d", message.ExamID, message.StudentID)

	// Check if exam exists
	exam, err := c.examRepo.FindByID(message.ExamID)
	if err != nil {
		log.Printf("Exam not found: %v", err)
		c.publisher.PublishSagaResponse(message.SagaItemID, false, "Exam not found")
		msg.Ack(false)
		return
	}

	// Check if student already registered
	existing, _ := c.examStudentRepo.FindByExamAndStudent(message.ExamID, message.StudentID)
	if existing != nil {
		log.Printf("Student already registered for exam")
		c.publisher.PublishSagaResponse(message.SagaItemID, false, "Student already registered")
		msg.Ack(false)
		return
	}

	// Check exam capacity
	count, _ := c.examStudentRepo.CountByExamID(message.ExamID)
	if int(count) >= exam.MaxStudentEntry {
		log.Printf("Exam is full")
		c.publisher.PublishSagaResponse(message.SagaItemID, false, "Exam capacity reached")
		msg.Ack(false)
		return
	}

	// Register student
	examStudent := &domain.ExamStudent{
		ExamID:    message.ExamID,
		StudentID: message.StudentID,
	}

	if err := c.examStudentRepo.Create(examStudent); err != nil {
		log.Printf("Failed to register student: %v", err)
		c.publisher.PublishSagaResponse(message.SagaItemID, false, err.Error())
		msg.Ack(false)
		return
	}

	log.Printf("Student successfully registered for exam")
	c.publisher.PublishSagaResponse(message.SagaItemID, true, "Student registered successfully")
	msg.Ack(false)
}

func (c *Consumer) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
}
