package domain

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

const (
	StatusInitialized = "initialized"
	StatusInProgress  = "in_progress"
	StatusFinished    = "finished"
)

const (
	SagaTypeUserClassVerification = "userClassVerificationSagaItem"
	SagaTypeUserWalletValidation  = "userWalletValidationSagaItem"
	SagaTypeUserWalletInsert      = "userWalletInsertSagaItem"
	SagaTypeExamRegistration      = "examRegistrationSagaItem"
)

type ExamRegistration struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	StudentID int        `json:"studentId" gorm:"column:student_id;not null"`
	ExamID    uint       `json:"examId" gorm:"column:exam_id;not null"`
	CourseID  int        `json:"courseId" gorm:"column:course_id;not null"`
	Status    string     `json:"status" gorm:"size:255;not null;default:initialized"`
	SagaItems []SagaItem `json:"sagaItems,omitempty" gorm:"foreignKey:ExamRegistrationID"`
	CreatedAt time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}

type SagaItem struct {
	ID                 uint             `json:"id" gorm:"primaryKey"`
	Type               string           `json:"type" gorm:"size:255;not null"` // validation or modification
	SagaType           string           `json:"sagaType" gorm:"column:saga_type;size:255;not null"`
	ExecutionOrder     int              `json:"executionOrder" gorm:"column:execution_order;not null"`
	Status             string           `json:"status" gorm:"size:255;not null;default:initialized"`
	StartedAt          *time.Time       `json:"startedAt" gorm:"column:started_at"`
	FinishedAt         *time.Time       `json:"finishedAt" gorm:"column:finished_at"`
	ReturnedPayload    JSON             `json:"returnedPayload" gorm:"column:returned_payload;type:json"`
	ExamRegistrationID uint             `json:"examRegistrationId" gorm:"column:exam_registration_id;not null"`
	ExamRegistration   *ExamRegistration `json:"examRegistration,omitempty" gorm:"foreignKey:ExamRegistrationID"`
	CreatedAt          time.Time        `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt          time.Time        `json:"updatedAt" gorm:"autoUpdateTime"`
}

// JSON type for GORM
type JSON map[string]interface{}

func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, j)
}

// SagaItemDefinition defines the saga steps
var SagaItemDefinitions = []struct {
	SagaType       string
	Type           string
	ExecutionOrder int
}{
	{SagaTypeUserClassVerification, "validation", 1},
	{SagaTypeUserWalletValidation, "validation", 2},
	{SagaTypeUserWalletInsert, "modification", 3},
	{SagaTypeExamRegistration, "modification", 4},
}

// Request DTOs
type CreateExamRegistrationRequest struct {
	StudentID int  `json:"studentId" binding:"required"`
	ExamID    uint `json:"examId" binding:"required"`
	CourseID  int  `json:"courseId" binding:"required"`
}
