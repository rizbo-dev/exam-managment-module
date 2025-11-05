package main

import (
	"gateway/internal/config"
	"gateway/internal/proxy"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	r := gin.Default()

	// CORS configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(corsConfig))

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Route definitions - proxy to respective services
	api := r.Group("/api")
	{
		// User service routes
		api.Any("/users/*path", proxy.ProxyRequest(cfg.UserServiceURL))

		// Exam service routes
		api.Any("/exams/*path", proxy.ProxyRequest(cfg.ExamServiceURL))
		api.Any("/examination-periods/*path", proxy.ProxyRequest(cfg.ExamServiceURL))

		// Class service routes
		api.Any("/students/*path", proxy.ProxyRequest(cfg.ClassServiceURL))
		api.Any("/departments/*path", proxy.ProxyRequest(cfg.ClassServiceURL))
		api.Any("/study-programs/*path", proxy.ProxyRequest(cfg.ClassServiceURL))
		api.Any("/courses/*path", proxy.ProxyRequest(cfg.ClassServiceURL))

		// Transaction service routes
		api.Any("/wallets/*path", proxy.ProxyRequest(cfg.TransactionServiceURL))
		api.Any("/transactions/*path", proxy.ProxyRequest(cfg.TransactionServiceURL))

		// Exam registration service routes
		api.Any("/exam-registrations/*path", proxy.ProxyRequest(cfg.ExamRegistrationServiceURL))
	}

	// Special endpoint for sign-for-exam (convenience endpoint)
	r.POST("/sign-for-exam", func(c *gin.Context) {
		// This proxies to exam-registration service
		proxy.ProxyRequest(cfg.ExamRegistrationServiceURL + "/api/exam-registrations")(c)
	})

	log.Printf("Gateway service starting on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
