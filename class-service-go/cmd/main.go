package main

import (
	"class-service/internal/config"
	"class-service/internal/domain"
	"class-service/internal/handler"
	"class-service/internal/messaging"
	"class-service/internal/repository"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	// Database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate
	if err := db.AutoMigrate(
		&domain.Department{},
		&domain.StudyProgram{},
		&domain.Course{},
		&domain.Student{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Setup repositories
	studentRepo := repository.NewStudentRepository(db)
	deptRepo := repository.NewDepartmentRepository(db)
	studyProgramRepo := repository.NewStudyProgramRepository(db)
	courseRepo := repository.NewCourseRepository(db)

	// Setup RabbitMQ
	publisher, err := messaging.NewPublisher(cfg.RabbitMQURL)
	if err != nil {
		log.Printf("Failed to setup RabbitMQ publisher: %v", err)
	}
	if publisher != nil {
		defer publisher.Close()
	}

	consumer, err := messaging.NewConsumer(cfg.RabbitMQURL, studentRepo, studyProgramRepo, courseRepo, publisher)
	if err != nil {
		log.Printf("Failed to setup RabbitMQ consumer: %v", err)
	} else {
		defer consumer.Close()
		if err := consumer.StartConsuming(); err != nil {
			log.Printf("Failed to start consuming: %v", err)
		}
	}

	// Setup handler
	classHandler := handler.NewClassHandler(studentRepo, deptRepo, studyProgramRepo, courseRepo)

	// Setup Gin router
	r := gin.Default()

	// CORS configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(corsConfig))

	// Routes
	api := r.Group("/api")
	{
		students := api.Group("/students")
		{
			students.POST("", classHandler.CreateStudent)
			students.GET("", classHandler.GetStudents)
			students.GET("/:id", classHandler.GetStudent)
		}

		departments := api.Group("/departments")
		{
			departments.POST("", classHandler.CreateDepartment)
			departments.GET("", classHandler.GetDepartments)
		}

		studyPrograms := api.Group("/study-programs")
		{
			studyPrograms.POST("", classHandler.CreateStudyProgram)
			studyPrograms.GET("", classHandler.GetStudyPrograms)
		}

		courses := api.Group("/courses")
		{
			courses.POST("", classHandler.CreateCourse)
			courses.GET("", classHandler.GetCourses)
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("Class service starting on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
