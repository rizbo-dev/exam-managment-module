package repository

import (
	"exam-registration-service/internal/domain"

	"gorm.io/gorm"
)

type ExamRegistrationRepository interface {
	Create(reg *domain.ExamRegistration) error
	FindByID(id uint) (*domain.ExamRegistration, error)
	FindAll(offset, limit int) ([]domain.ExamRegistration, error)
	Update(reg *domain.ExamRegistration) error
	Count() (int64, error)
}

type examRegistrationRepository struct {
	db *gorm.DB
}

func NewExamRegistrationRepository(db *gorm.DB) ExamRegistrationRepository {
	return &examRegistrationRepository{db: db}
}

func (r *examRegistrationRepository) Create(reg *domain.ExamRegistration) error {
	return r.db.Create(reg).Error
}

func (r *examRegistrationRepository) FindByID(id uint) (*domain.ExamRegistration, error) {
	var reg domain.ExamRegistration
	err := r.db.Preload("SagaItems").First(&reg, id).Error
	if err != nil {
		return nil, err
	}
	return &reg, nil
}

func (r *examRegistrationRepository) FindAll(offset, limit int) ([]domain.ExamRegistration, error) {
	var regs []domain.ExamRegistration
	err := r.db.Preload("SagaItems").Offset(offset).Limit(limit).Find(&regs).Error
	return regs, err
}

func (r *examRegistrationRepository) Update(reg *domain.ExamRegistration) error {
	return r.db.Save(reg).Error
}

func (r *examRegistrationRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&domain.ExamRegistration{}).Count(&count).Error
	return count, err
}

// SagaItem Repository
type SagaItemRepository interface {
	Create(item *domain.SagaItem) error
	FindByID(id uint) (*domain.SagaItem, error)
	Update(item *domain.SagaItem) error
	FindByExamRegistrationID(regID uint) ([]domain.SagaItem, error)
}

type sagaItemRepository struct {
	db *gorm.DB
}

func NewSagaItemRepository(db *gorm.DB) SagaItemRepository {
	return &sagaItemRepository{db: db}
}

func (r *sagaItemRepository) Create(item *domain.SagaItem) error {
	return r.db.Create(item).Error
}

func (r *sagaItemRepository) FindByID(id uint) (*domain.SagaItem, error) {
	var item domain.SagaItem
	err := r.db.Preload("ExamRegistration").First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *sagaItemRepository) Update(item *domain.SagaItem) error {
	return r.db.Save(item).Error
}

func (r *sagaItemRepository) FindByExamRegistrationID(regID uint) ([]domain.SagaItem, error) {
	var items []domain.SagaItem
	err := r.db.Where("exam_registration_id = ?", regID).Order("execution_order ASC").Find(&items).Error
	return items, err
}
