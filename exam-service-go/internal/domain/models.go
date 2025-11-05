package domain

import "time"

type ExaminationPeriod struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	StartDate time.Time `json:"startDate" gorm:"column:start_date;not null"`
	EndDate   time.Time `json:"endDate" gorm:"column:end_date;not null"`
	Exams     []Exam    `json:"exams,omitempty" gorm:"foreignKey:ExaminationPeriodID"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type Exam struct {
	ID                   uint               `json:"id" gorm:"primaryKey"`
	CourseID             int                `json:"courseId" gorm:"column:course_id;not null"`
	ExaminationPeriodID  *uint              `json:"examinationPeriodId" gorm:"column:examination_period_id"`
	ExaminationPeriod    *ExaminationPeriod `json:"examinationPeriod,omitempty" gorm:"foreignKey:ExaminationPeriodID"`
	MaxStudentEntry      int                `json:"maxStudentEntry" gorm:"column:max_student_entry;not null"`
	Cost                 float64            `json:"cost" gorm:"not null"`
	ExamStudents         []ExamStudent      `json:"examStudents,omitempty" gorm:"foreignKey:ExamID"`
	CreatedAt            time.Time          `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt            time.Time          `json:"updatedAt" gorm:"autoUpdateTime"`
}

type ExamStudent struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ExamID    uint      `json:"examId" gorm:"column:exam_id;not null;uniqueIndex:idx_exam_student"`
	Exam      *Exam     `json:"exam,omitempty" gorm:"foreignKey:ExamID"`
	StudentID int       `json:"studentId" gorm:"column:student_id;not null;uniqueIndex:idx_exam_student"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type ExamResult struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ExamID    uint      `json:"examId" gorm:"column:exam_id;not null"`
	Exam      *Exam     `json:"exam,omitempty" gorm:"foreignKey:ExamID"`
	StudentID int       `json:"studentId" gorm:"column:student_id;not null"`
	Score     float64   `json:"score" gorm:"not null"`
	Passed    bool      `json:"passed" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// Request/Response DTOs
type CreateExamRequest struct {
	CourseID            int     `json:"courseId" binding:"required"`
	ExaminationPeriodID *uint   `json:"examinationPeriodId"`
	MaxStudentEntry     int     `json:"maxStudentEntry" binding:"required"`
	Cost                float64 `json:"cost" binding:"required"`
}

type CreateExaminationPeriodRequest struct {
	StartDate time.Time `json:"startDate" binding:"required"`
	EndDate   time.Time `json:"endDate" binding:"required"`
}

type RegisterStudentRequest struct {
	ExamID    uint `json:"examId" binding:"required"`
	StudentID int  `json:"studentId" binding:"required"`
}
