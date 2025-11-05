package handler

import (
	"class-service/internal/domain"
	"class-service/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClassHandler struct {
	studentRepo      repository.StudentRepository
	deptRepo         repository.DepartmentRepository
	studyProgramRepo repository.StudyProgramRepository
	courseRepo       repository.CourseRepository
}

func NewClassHandler(
	studentRepo repository.StudentRepository,
	deptRepo repository.DepartmentRepository,
	studyProgramRepo repository.StudyProgramRepository,
	courseRepo repository.CourseRepository,
) *ClassHandler {
	return &ClassHandler{
		studentRepo:      studentRepo,
		deptRepo:         deptRepo,
		studyProgramRepo: studyProgramRepo,
		courseRepo:       courseRepo,
	}
}

// Student handlers
func (h *ClassHandler) CreateStudent(c *gin.Context) {
	var req domain.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := &domain.Student{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		DepartmentID:   req.DepartmentID,
		StudyProgramID: req.StudyProgramID,
	}

	if err := h.studentRepo.Create(student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

func (h *ClassHandler) GetStudent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	student, err := h.studentRepo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *ClassHandler) GetStudents(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "30"))

	offset := (page - 1) * limit

	students, err := h.studentRepo.FindAll(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, _ := h.studentRepo.Count()

	c.JSON(http.StatusOK, gin.H{
		"data":       students,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}

// Department handlers
func (h *ClassHandler) CreateDepartment(c *gin.Context) {
	var req domain.CreateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dept := &domain.Department{Name: req.Name}
	if err := h.deptRepo.Create(dept); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dept)
}

func (h *ClassHandler) GetDepartments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "30"))

	offset := (page - 1) * limit

	depts, err := h.deptRepo.FindAll(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, _ := h.deptRepo.Count()

	c.JSON(http.StatusOK, gin.H{
		"data":       depts,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}

// StudyProgram handlers
func (h *ClassHandler) CreateStudyProgram(c *gin.Context) {
	var req domain.CreateStudyProgramRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	program := &domain.StudyProgram{
		Name:         req.Name,
		DepartmentID: req.DepartmentID,
	}

	if err := h.studyProgramRepo.Create(program); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, program)
}

func (h *ClassHandler) GetStudyPrograms(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "30"))

	offset := (page - 1) * limit

	programs, err := h.studyProgramRepo.FindAll(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, _ := h.studyProgramRepo.Count()

	c.JSON(http.StatusOK, gin.H{
		"data":       programs,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}

// Course handlers
func (h *ClassHandler) CreateCourse(c *gin.Context) {
	var req domain.CreateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course := &domain.Course{
		Name:      req.Name,
		Code:      req.Code,
		TeacherID: req.TeacherID,
	}

	if err := h.courseRepo.Create(course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, course)
}

func (h *ClassHandler) GetCourses(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "30"))

	offset := (page - 1) * limit

	courses, err := h.courseRepo.FindAll(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, _ := h.courseRepo.Count()

	c.JSON(http.StatusOK, gin.H{
		"data":       courses,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}
