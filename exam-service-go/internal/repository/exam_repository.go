package repository

import (
	"exam-service/internal/domain"

	"gorm.io/gorm"
)

type ExamRepository interface {
	Create(exam *domain.Exam) error
	FindByID(id uint) (*domain.Exam, error)
	FindAll(offset, limit int) ([]domain.Exam, error)
	Update(exam *domain.Exam) error
	Delete(id uint) error
	Count() (int64, error)
}

type examRepository struct {
	db *gorm.DB
}

func NewExamRepository(db *gorm.DB) ExamRepository {
	return &examRepository{db: db}
}

func (r *examRepository) Create(exam *domain.Exam) error {
	return r.db.Create(exam).Error
}

func (r *examRepository) FindByID(id uint) (*domain.Exam, error) {
	var exam domain.Exam
	err := r.db.Preload("ExaminationPeriod").Preload("ExamStudents").First(&exam, id).Error
	if err != nil {
		return nil, err
	}
	return &exam, nil
}

func (r *examRepository) FindAll(offset, limit int) ([]domain.Exam, error) {
	var exams []domain.Exam
	err := r.db.Preload("ExaminationPeriod").Offset(offset).Limit(limit).Find(&exams).Error
	return exams, err
}

func (r *examRepository) Update(exam *domain.Exam) error {
	return r.db.Save(exam).Error
}

func (r *examRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Exam{}, id).Error
}

func (r *examRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&domain.Exam{}).Count(&count).Error
	return count, err
}
