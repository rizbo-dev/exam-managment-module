package domain

import "time"

type User struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	FirstName      string    `json:"firstName" gorm:"column:first_name;size:255;not null"`
	LastName       string    `json:"lastName" gorm:"column:last_name;size:255;not null"`
	Role           string    `json:"role" gorm:"size:255;not null"`
	DepartmentID   int       `json:"departmentId" gorm:"column:department_id;not null"`
	StudyProgramID int       `json:"studyProgramId" gorm:"column:study_program_id;not null"`
	CreatedAt      time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "user"
}

type CreateUserRequest struct {
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`
	Role           string `json:"role" binding:"required"`
	DepartmentID   int    `json:"departmentId" binding:"required"`
	StudyProgramID int    `json:"studyProgramId" binding:"required"`
}

type UpdateUserRequest struct {
	FirstName      *string `json:"firstName,omitempty"`
	LastName       *string `json:"lastName,omitempty"`
	Role           *string `json:"role,omitempty"`
	DepartmentID   *int    `json:"departmentId,omitempty"`
	StudyProgramID *int    `json:"studyProgramId,omitempty"`
}
