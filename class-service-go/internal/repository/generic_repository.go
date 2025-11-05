package repository

import (
	"class-service/internal/domain"

	"gorm.io/gorm"
)

// Department Repository
type DepartmentRepository interface {
	Create(dept *domain.Department) error
	FindByID(id uint) (*domain.Department, error)
	FindAll(offset, limit int) ([]domain.Department, error)
	Update(dept *domain.Department) error
	Delete(id uint) error
	Count() (int64, error)
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{db: db}
}

func (r *departmentRepository) Create(dept *domain.Department) error {
	return r.db.Create(dept).Error
}

func (r *departmentRepository) FindByID(id uint) (*domain.Department, error) {
	var dept domain.Department
	err := r.db.Preload("StudyPrograms").First(&dept, id).Error
	if err != nil {
		return nil, err
	}
	return &dept, nil
}

func (r *departmentRepository) FindAll(offset, limit int) ([]domain.Department, error) {
	var depts []domain.Department
	err := r.db.Offset(offset).Limit(limit).Find(&depts).Error
	return depts, err
}

func (r *departmentRepository) Update(dept *domain.Department) error {
	return r.db.Save(dept).Error
}

func (r *departmentRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Department{}, id).Error
}

func (r *departmentRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&domain.Department{}).Count(&count).Error
	return count, err
}

// StudyProgram Repository
type StudyProgramRepository interface {
	Create(program *domain.StudyProgram) error
	FindByID(id uint) (*domain.StudyProgram, error)
	FindAll(offset, limit int) ([]domain.StudyProgram, error)
	Update(program *domain.StudyProgram) error
	Delete(id uint) error
	Count() (int64, error)
}

type studyProgramRepository struct {
	db *gorm.DB
}

func NewStudyProgramRepository(db *gorm.DB) StudyProgramRepository {
	return &studyProgramRepository{db: db}
}

func (r *studyProgramRepository) Create(program *domain.StudyProgram) error {
	return r.db.Create(program).Error
}

func (r *studyProgramRepository) FindByID(id uint) (*domain.StudyProgram, error) {
	var program domain.StudyProgram
	err := r.db.Preload("Department").Preload("Courses").First(&program, id).Error
	if err != nil {
		return nil, err
	}
	return &program, nil
}

func (r *studyProgramRepository) FindAll(offset, limit int) ([]domain.StudyProgram, error) {
	var programs []domain.StudyProgram
	err := r.db.Preload("Department").Offset(offset).Limit(limit).Find(&programs).Error
	return programs, err
}

func (r *studyProgramRepository) Update(program *domain.StudyProgram) error {
	return r.db.Save(program).Error
}

func (r *studyProgramRepository) Delete(id uint) error {
	return r.db.Delete(&domain.StudyProgram{}, id).Error
}

func (r *studyProgramRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&domain.StudyProgram{}).Count(&count).Error
	return count, err
}

// Course Repository
type CourseRepository interface {
	Create(course *domain.Course) error
	FindByID(id uint) (*domain.Course, error)
	FindAll(offset, limit int) ([]domain.Course, error)
	Update(course *domain.Course) error
	Delete(id uint) error
	Count() (int64, error)
	FindByStudyProgramID(studyProgramID uint) ([]domain.Course, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) Create(course *domain.Course) error {
	return r.db.Create(course).Error
}

func (r *courseRepository) FindByID(id uint) (*domain.Course, error) {
	var course domain.Course
	err := r.db.Preload("StudyPrograms").First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *courseRepository) FindAll(offset, limit int) ([]domain.Course, error) {
	var courses []domain.Course
	err := r.db.Offset(offset).Limit(limit).Find(&courses).Error
	return courses, err
}

func (r *courseRepository) Update(course *domain.Course) error {
	return r.db.Save(course).Error
}

func (r *courseRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Course{}, id).Error
}

func (r *courseRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&domain.Course{}).Count(&count).Error
	return count, err
}

func (r *courseRepository) FindByStudyProgramID(studyProgramID uint) ([]domain.Course, error) {
	var studyProgram domain.StudyProgram
	err := r.db.Preload("Courses").First(&studyProgram, studyProgramID).Error
	if err != nil {
		return nil, err
	}
	return studyProgram.Courses, nil
}
