package repository

import (
	"class-service/internal/domain"

	"gorm.io/gorm"
)

type StudentRepository interface {
	Create(student *domain.Student) error
	FindByID(id uint) (*domain.Student, error)
	FindAll(offset, limit int) ([]domain.Student, error)
	Update(student *domain.Student) error
	Delete(id uint) error
	Count() (int64, error)
	FindByStudyProgramID(studyProgramID uint) ([]domain.Student, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (r *studentRepository) Create(student *domain.Student) error {
	return r.db.Create(student).Error
}

func (r *studentRepository) FindByID(id uint) (*domain.Student, error) {
	var student domain.Student
	err := r.db.Preload("Department").Preload("StudyProgram").First(&student, id).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepository) FindAll(offset, limit int) ([]domain.Student, error) {
	var students []domain.Student
	err := r.db.Preload("Department").Preload("StudyProgram").Offset(offset).Limit(limit).Find(&students).Error
	return students, err
}

func (r *studentRepository) Update(student *domain.Student) error {
	return r.db.Save(student).Error
}

func (r *studentRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Student{}, id).Error
}

func (r *studentRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&domain.Student{}).Count(&count).Error
	return count, err
}

func (r *studentRepository) FindByStudyProgramID(studyProgramID uint) ([]domain.Student, error) {
	var students []domain.Student
	err := r.db.Where("study_program_id = ?", studyProgramID).Find(&students).Error
	return students, err
}
