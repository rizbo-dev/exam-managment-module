package main

import (
	"exam-registration-service/internal/config"
	"exam-registration-service/internal/domain"
	"exam-registration-service/internal/handler"
	"exam-registration-service/internal/messaging"
	"exam-registration-service/internal/repository"
	"exam-registration-service/internal/service"
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
		&domain.ExamRegistration{},
		&domain.SagaItem{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Setup repositories
	regRepo := repository.NewExamRegistrationRepository(db)
	itemRepo := repository.NewSagaItemRepository(db)

	// Setup RabbitMQ publisher
	publisher, err := messaging.NewPublisher(cfg.RabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to setup RabbitMQ publisher: %v", err)
	}
	defer publisher.Close()

	// Setup saga orchestrator
	orchestrator := service.NewSagaOrchestrator(regRepo, itemRepo, publisher)

	// Setup RabbitMQ consumer
	consumer, err := messaging.NewConsumer(cfg.RabbitMQURL, orchestrator)
	if err != nil {
		log.Fatalf("Failed to setup RabbitMQ consumer: %v", err)
	}
	defer consumer.Close()

	if err := consumer.StartConsuming(); err != nil {
		log.Fatalf("Failed to start consuming: %v", err)
	}

	// Setup handler
	regHandler := handler.NewExamRegistrationHandler(regRepo, orchestrator)

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
		registrations := api.Group("/exam-registrations")
		{
			registrations.POST("", regHandler.CreateExamRegistration)
			registrations.GET("", regHandler.GetExamRegistrations)
			registrations.GET("/:id", regHandler.GetExamRegistration)
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("Exam registration service starting on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
