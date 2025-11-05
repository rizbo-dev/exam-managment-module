package main

import (
	"fmt"
	"log"
	"transaction-service/internal/config"
	"transaction-service/internal/domain"
	"transaction-service/internal/handler"
	"transaction-service/internal/messaging"
	"transaction-service/internal/repository"

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
		&domain.Wallet{},
		&domain.Transaction{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Setup repositories
	walletRepo := repository.NewWalletRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	// Setup RabbitMQ
	publisher, err := messaging.NewPublisher(cfg.RabbitMQURL)
	if err != nil {
		log.Printf("Failed to setup RabbitMQ publisher: %v", err)
	}
	if publisher != nil {
		defer publisher.Close()
	}

	consumer, err := messaging.NewConsumer(cfg.RabbitMQURL, walletRepo, transactionRepo, publisher)
	if err != nil {
		log.Printf("Failed to setup RabbitMQ consumer: %v", err)
	} else {
		defer consumer.Close()
		if err := consumer.StartConsuming(); err != nil {
			log.Printf("Failed to start consuming: %v", err)
		}
	}

	// Setup handler
	transactionHandler := handler.NewTransactionHandler(walletRepo, transactionRepo)

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
		wallets := api.Group("/wallets")
		{
			wallets.POST("", transactionHandler.CreateWallet)
			wallets.GET("", transactionHandler.GetWallets)
			wallets.GET("/student/:studentId", transactionHandler.GetWalletByStudentID)
			wallets.POST("/student/:studentId/deposit", transactionHandler.Deposit)
			wallets.GET("/student/:studentId/transactions", transactionHandler.GetWalletTransactions)
		}

		transactions := api.Group("/transactions")
		{
			transactions.GET("", transactionHandler.GetTransactions)
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("Transaction service starting on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
