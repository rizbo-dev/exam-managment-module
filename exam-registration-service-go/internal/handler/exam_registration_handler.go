package handler

import (
	"exam-registration-service/internal/domain"
	"exam-registration-service/internal/repository"
	"exam-registration-service/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExamRegistrationHandler struct {
	repo         repository.ExamRegistrationRepository
	orchestrator *service.SagaOrchestrator
}

func NewExamRegistrationHandler(
	repo repository.ExamRegistrationRepository,
	orchestrator *service.SagaOrchestrator,
) *ExamRegistrationHandler {
	return &ExamRegistrationHandler{
		repo:         repo,
		orchestrator: orchestrator,
	}
}

func (h *ExamRegistrationHandler) CreateExamRegistration(c *gin.Context) {
	var req domain.CreateExamRegistrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registration, err := h.orchestrator.StartExamRegistration(req.StudentID, req.ExamID, req.CourseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, registration)
}

func (h *ExamRegistrationHandler) GetExamRegistration(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid registration ID"})
		return
	}

	registration, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Registration not found"})
		return
	}

	c.JSON(http.StatusOK, registration)
}

func (h *ExamRegistrationHandler) GetExamRegistrations(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "30"))

	offset := (page - 1) * limit

	registrations, err := h.repo.FindAll(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, _ := h.repo.Count()

	c.JSON(http.StatusOK, gin.H{
		"data":       registrations,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}
