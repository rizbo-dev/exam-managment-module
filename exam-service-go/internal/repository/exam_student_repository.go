package repository

import (
	"exam-service/internal/domain"

	"gorm.io/gorm"
)

type ExamStudentRepository interface {
	Create(examStudent *domain.ExamStudent) error
	FindByExamAndStudent(examID uint, studentID int) (*domain.ExamStudent, error)
	FindByExamID(examID uint) ([]domain.ExamStudent, error)
	FindByStudentID(studentID int) ([]domain.ExamStudent, error)
	Delete(id uint) error
	CountByExamID(examID uint) (int64, error)
}

type examStudentRepository struct {
	db *gorm.DB
}

func NewExamStudentRepository(db *gorm.DB) ExamStudentRepository {
	return &examStudentRepository{db: db}
}

func (r *examStudentRepository) Create(examStudent *domain.ExamStudent) error {
	return r.db.Create(examStudent).Error
}

func (r *examStudentRepository) FindByExamAndStudent(examID uint, studentID int) (*domain.ExamStudent, error) {
	var examStudent domain.ExamStudent
	err := r.db.Where("exam_id = ? AND student_id = ?", examID, studentID).First(&examStudent).Error
	if err != nil {
		return nil, err
	}
	return &examStudent, nil
}

func (r *examStudentRepository) FindByExamID(examID uint) ([]domain.ExamStudent, error) {
	var examStudents []domain.ExamStudent
	err := r.db.Where("exam_id = ?", examID).Find(&examStudents).Error
	return examStudents, err
}

func (r *examStudentRepository) FindByStudentID(studentID int) ([]domain.ExamStudent, error) {
	var examStudents []domain.ExamStudent
	err := r.db.Preload("Exam").Where("student_id = ?", studentID).Find(&examStudents).Error
	return examStudents, err
}

func (r *examStudentRepository) Delete(id uint) error {
	return r.db.Delete(&domain.ExamStudent{}, id).Error
}

func (r *examStudentRepository) CountByExamID(examID uint) (int64, error) {
	var count int64
	err := r.db.Model(&domain.ExamStudent{}).Where("exam_id = ?", examID).Count(&count).Error
	return count, err
}
