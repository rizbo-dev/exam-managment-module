package repository

import (
	"exam-service/internal/domain"

	"gorm.io/gorm"
)

type ExaminationPeriodRepository interface {
	Create(period *domain.ExaminationPeriod) error
	FindByID(id uint) (*domain.ExaminationPeriod, error)
	FindAll(offset, limit int) ([]domain.ExaminationPeriod, error)
	Update(period *domain.ExaminationPeriod) error
	Delete(id uint) error
	Count() (int64, error)
}

type examinationPeriodRepository struct {
	db *gorm.DB
}

func NewExaminationPeriodRepository(db *gorm.DB) ExaminationPeriodRepository {
	return &examinationPeriodRepository{db: db}
}

func (r *examinationPeriodRepository) Create(period *domain.ExaminationPeriod) error {
	return r.db.Create(period).Error
}

func (r *examinationPeriodRepository) FindByID(id uint) (*domain.ExaminationPeriod, error) {
	var period domain.ExaminationPeriod
	err := r.db.Preload("Exams").First(&period, id).Error
	if err != nil {
		return nil, err
	}
	return &period, nil
}

func (r *examinationPeriodRepository) FindAll(offset, limit int) ([]domain.ExaminationPeriod, error) {
	var periods []domain.ExaminationPeriod
	err := r.db.Offset(offset).Limit(limit).Find(&periods).Error
	return periods, err
}

func (r *examinationPeriodRepository) Update(period *domain.ExaminationPeriod) error {
	return r.db.Save(period).Error
}

func (r *examinationPeriodRepository) Delete(id uint) error {
	return r.db.Delete(&domain.ExaminationPeriod{}, id).Error
}

func (r *examinationPeriodRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&domain.ExaminationPeriod{}).Count(&count).Error
	return count, err
}
