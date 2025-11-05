package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort                 string
	UserServiceURL             string
	ExamServiceURL             string
	ClassServiceURL            string
	TransactionServiceURL      string
	EmailServiceURL            string
	ExamRegistrationServiceURL string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		ServerPort:                 getEnv("SERVER_PORT", "8086"),
		UserServiceURL:             getEnv("USER_SERVICE_URL", "http://user-service-go:8080"),
		ExamServiceURL:             getEnv("EXAM_SERVICE_URL", "http://exam-service-go:8081"),
		ClassServiceURL:            getEnv("CLASS_SERVICE_URL", "http://class-service-go:8082"),
		TransactionServiceURL:      getEnv("TRANSACTION_SERVICE_URL", "http://transaction-service-go:8083"),
		EmailServiceURL:            getEnv("EMAIL_SERVICE_URL", "http://email-service-go:8084"),
		ExamRegistrationServiceURL: getEnv("EXAM_REGISTRATION_SERVICE_URL", "http://exam-registration-service-go:8085"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
