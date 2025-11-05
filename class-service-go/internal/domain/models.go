package domain

import "time"

type Department struct {
	ID            uint            `json:"id" gorm:"primaryKey"`
	Name          string          `json:"name" gorm:"size:255;not null"`
	StudyPrograms []StudyProgram  `json:"studyPrograms,omitempty" gorm:"foreignKey:DepartmentID"`
	Students      []Student       `json:"students,omitempty" gorm:"foreignKey:DepartmentID"`
	CreatedAt     time.Time       `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time       `json:"updatedAt" gorm:"autoUpdateTime"`
}

type StudyProgram struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"size:255;not null"`
	DepartmentID uint      `json:"departmentId" gorm:"column:department_id;not null"`
	Department   *Department `json:"department,omitempty" gorm:"foreignKey:DepartmentID"`
	Students     []Student `json:"students,omitempty" gorm:"foreignKey:StudyProgramID"`
	Courses      []Course  `json:"courses,omitempty" gorm:"many2many:study_program_courses;"`
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type Course struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Name          string         `json:"name" gorm:"size:255;not null"`
	Code          string         `json:"code" gorm:"size:50;not null;unique"`
	TeacherID     *int           `json:"teacherId" gorm:"column:teacher_id"`
	StudyPrograms []StudyProgram `json:"studyPrograms,omitempty" gorm:"many2many:study_program_courses;"`
	CreatedAt     time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
}

type Student struct {
	ID             uint          `json:"id" gorm:"primaryKey"`
	FirstName      string        `json:"firstName" gorm:"column:first_name;size:255;not null"`
	LastName       string        `json:"lastName" gorm:"column:last_name;size:255;not null"`
	Email          string        `json:"email" gorm:"size:255;not null;unique"`
	DepartmentID   uint          `json:"departmentId" gorm:"column:department_id;not null"`
	Department     *Department   `json:"department,omitempty" gorm:"foreignKey:DepartmentID"`
	StudyProgramID uint          `json:"studyProgramId" gorm:"column:study_program_id;not null"`
	StudyProgram   *StudyProgram `json:"studyProgram,omitempty" gorm:"foreignKey:StudyProgramID"`
	CreatedAt      time.Time     `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      time.Time     `json:"updatedAt" gorm:"autoUpdateTime"`
}

// Request DTOs
type CreateStudentRequest struct {
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	DepartmentID   uint   `json:"departmentId" binding:"required"`
	StudyProgramID uint   `json:"studyProgramId" binding:"required"`
}

type CreateDepartmentRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateStudyProgramRequest struct {
	Name         string `json:"name" binding:"required"`
	DepartmentID uint   `json:"departmentId" binding:"required"`
}

type CreateCourseRequest struct {
	Name      string `json:"name" binding:"required"`
	Code      string `json:"code" binding:"required"`
	TeacherID *int   `json:"teacherId"`
}
