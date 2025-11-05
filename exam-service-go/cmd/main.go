package main

import (
	"fmt"
	"log"
	"exam-service/internal/config"
	"exam-service/internal/domain"
	"exam-service/internal/handler"
	"exam-service/internal/messaging"
	"exam-service/internal/repository"

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
		&domain.ExaminationPeriod{},
		&domain.Exam{},
		&domain.ExamStudent{},
		&domain.ExamResult{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Setup repositories
	examRepo := repository.NewExamRepository(db)
	periodRepo := repository.NewExaminationPeriodRepository(db)
	examStudentRepo := repository.NewExamStudentRepository(db)

	// Setup RabbitMQ
	publisher, err := messaging.NewPublisher(cfg.RabbitMQURL)
	if err != nil {
		log.Printf("Failed to setup RabbitMQ publisher: %v", err)
	}
	if publisher != nil {
		defer publisher.Close()
	}

	consumer, err := messaging.NewConsumer(cfg.RabbitMQURL, examRepo, examStudentRepo, publisher)
	if err != nil {
		log.Printf("Failed to setup RabbitMQ consumer: %v", err)
	} else {
		defer consumer.Close()
		if err := consumer.StartConsuming(); err != nil {
			log.Printf("Failed to start consuming: %v", err)
		}
	}

	// Setup handler
	examHandler := handler.NewExamHandler(examRepo, periodRepo, examStudentRepo)

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
		exams := api.Group("/exams")
		{
			exams.POST("", examHandler.CreateExam)
			exams.GET("", examHandler.GetExams)
			exams.GET("/:id", examHandler.GetExam)
			exams.DELETE("/:id", examHandler.DeleteExam)
		}

		periods := api.Group("/examination-periods")
		{
			periods.POST("", examHandler.CreateExaminationPeriod)
			periods.GET("", examHandler.GetExaminationPeriods)
			periods.GET("/:id", examHandler.GetExaminationPeriod)
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("Exam service starting on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
