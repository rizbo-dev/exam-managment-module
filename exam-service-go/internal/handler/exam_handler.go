package handler

import (
	"exam-service/internal/domain"
	"exam-service/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExamHandler struct {
	examRepo        repository.ExamRepository
	periodRepo      repository.ExaminationPeriodRepository
	examStudentRepo repository.ExamStudentRepository
}

func NewExamHandler(
	examRepo repository.ExamRepository,
	periodRepo repository.ExaminationPeriodRepository,
	examStudentRepo repository.ExamStudentRepository,
) *ExamHandler {
	return &ExamHandler{
		examRepo:        examRepo,
		periodRepo:      periodRepo,
		examStudentRepo: examStudentRepo,
	}
}

// Exam endpoints
func (h *ExamHandler) CreateExam(c *gin.Context) {
	var req domain.CreateExamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exam := &domain.Exam{
		CourseID:            req.CourseID,
		ExaminationPeriodID: req.ExaminationPeriodID,
		MaxStudentEntry:     req.MaxStudentEntry,
		Cost:                req.Cost,
	}

	if err := h.examRepo.Create(exam); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, exam)
}

func (h *ExamHandler) GetExam(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam ID"})
		return
	}

	exam, err := h.examRepo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exam not found"})
		return
	}

	c.JSON(http.StatusOK, exam)
}

func (h *ExamHandler) GetExams(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "30"))

	offset := (page - 1) * limit

	exams, err := h.examRepo.FindAll(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, _ := h.examRepo.Count()

	c.JSON(http.StatusOK, gin.H{
		"data":       exams,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}

func (h *ExamHandler) DeleteExam(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam ID"})
		return
	}

	if err := h.examRepo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Examination Period endpoints
func (h *ExamHandler) CreateExaminationPeriod(c *gin.Context) {
	var req domain.CreateExaminationPeriodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	period := &domain.ExaminationPeriod{
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}

	if err := h.periodRepo.Create(period); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, period)
}

func (h *ExamHandler) GetExaminationPeriod(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid period ID"})
		return
	}

	period, err := h.periodRepo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Period not found"})
		return
	}

	c.JSON(http.StatusOK, period)
}

func (h *ExamHandler) GetExaminationPeriods(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "30"))

	offset := (page - 1) * limit

	periods, err := h.periodRepo.FindAll(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, _ := h.periodRepo.Count()

	c.JSON(http.StatusOK, gin.H{
		"data":       periods,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}
