# Student Exam Management System - Go Edition

This is a complete rewrite of the Student Exam Management System microservices from PHP/Symfony to Go (Golang).

## Architecture Overview

The system consists of 7 microservices built with Go:

| Service | Port | Description |
|---------|------|-------------|
| **user-service** | 8080 | User profile management |
| **exam-service** | 8081 | Exam definitions & examination periods |
| **class-service** | 8082 | Students, courses, departments, and study programs |
| **transaction-service** | 8083 | Wallet & payment processing |
| **email-service** | 8084 | Email notifications |
| **exam-registration-service** | 8085 | Exam registration saga orchestration |
| **gateway** | 8086 | API Gateway & request routing |

## Technology Stack

- **Language:** Go 1.21+
- **Web Framework:** Gin
- **ORM:** GORM
- **Message Queue:** RabbitMQ (amqp091-go)
- **Database:** MySQL 8
- **Containerization:** Docker & Docker Compose

## Key Features

### 1. Saga Pattern Implementation
The exam registration process uses the Saga pattern with 4 orchestrated steps:
1. **User Class Verification** - Validates student enrollment in course
2. **User Wallet Validation** - Checks sufficient funds
3. **User Wallet Insert** - Deducts exam fee
4. **Exam Registration** - Registers student for exam

### 2. Event-Driven Architecture
Services communicate asynchronously via RabbitMQ exchanges:
- `user-class-verification` - Class enrollment verification
- `user-wallet-validation` - Wallet balance validation
- `user-wallet-insert` - Wallet deduction
- `exam-registration` - Final exam registration
- `response-saga-item` - Saga step responses

### 3. Clean Architecture
Each service follows clean architecture principles:
```
service-name-go/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── domain/              # Domain models & business entities
│   ├── repository/          # Data access layer
│   ├── handler/             # HTTP handlers
│   ├── messaging/           # RabbitMQ publishers/consumers
│   ├── service/             # Business logic
│   └── config/              # Configuration management
├── go.mod
├── go.sum
├── Dockerfile
└── .env
```

## Prerequisites

- Docker & Docker Compose
- Go 1.21+ (for local development)
- MySQL 8
- RabbitMQ 3

## Getting Started

### Option 1: Using Docker Compose (Recommended)

1. **Clone the repository:**
```bash
git clone <repository-url>
cd exam-managment-module
```

2. **Build and start all services:**
```bash
docker-compose -f docker-compose-go.yml up --build
```

3. **Wait for services to be ready:**
   - MySQL healthcheck will ensure database is ready
   - RabbitMQ management UI: http://localhost:15672 (guest/guest)
   - Mailhog UI: http://localhost:8025

4. **Access services:**
   - Gateway: http://localhost:8086
   - User Service: http://localhost:8080
   - Exam Service: http://localhost:8081
   - Class Service: http://localhost:8082
   - Transaction Service: http://localhost:8083
   - Email Service: http://localhost:8084
   - Exam Registration Service: http://localhost:8085

### Option 2: Local Development

Each service can be run locally for development:

```bash
# Example: Running user-service locally
cd user-service-go

# Install dependencies
go mod download

# Run the service
go run cmd/main.go
```

**Note:** Update the `.env` file in each service to point to your local MySQL and RabbitMQ instances.

## API Endpoints

### Gateway Routes (Port 8086)
All routes are prefixed with `/api`:

**Users:**
- `GET /api/users` - List users
- `POST /api/users` - Create user
- `GET /api/users/:id` - Get user
- `PUT /api/users/:id` - Update user
- `DELETE /api/users/:id` - Delete user

**Exams:**
- `GET /api/exams` - List exams
- `POST /api/exams` - Create exam
- `GET /api/exams/:id` - Get exam
- `DELETE /api/exams/:id` - Delete exam

**Examination Periods:**
- `GET /api/examination-periods` - List periods
- `POST /api/examination-periods` - Create period
- `GET /api/examination-periods/:id` - Get period

**Students:**
- `GET /api/students` - List students
- `POST /api/students` - Create student
- `GET /api/students/:id` - Get student

**Departments & Study Programs:**
- `GET /api/departments` - List departments
- `POST /api/departments` - Create department
- `GET /api/study-programs` - List study programs
- `POST /api/study-programs` - Create study program

**Courses:**
- `GET /api/courses` - List courses
- `POST /api/courses` - Create course

**Wallets:**
- `GET /api/wallets` - List wallets
- `POST /api/wallets` - Create wallet
- `GET /api/wallets/student/:studentId` - Get wallet by student
- `POST /api/wallets/student/:studentId/deposit` - Deposit to wallet
- `GET /api/wallets/student/:studentId/transactions` - Get transactions

**Exam Registrations:**
- `POST /api/exam-registrations` - Start exam registration (triggers saga)
- `GET /api/exam-registrations` - List registrations
- `GET /api/exam-registrations/:id` - Get registration status

**Special Endpoints:**
- `POST /sign-for-exam` - Convenience endpoint for exam registration

## Example: Exam Registration Flow

```bash
# 1. Create a wallet for student
curl -X POST http://localhost:8086/api/wallets \
  -H "Content-Type: application/json" \
  -d '{
    "studentId": 1,
    "balance": 100.00
  }'

# 2. Deposit funds
curl -X POST http://localhost:8086/api/wallets/student/1/deposit \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 50.00
  }'

# 3. Register for exam (triggers saga)
curl -X POST http://localhost:8086/api/exam-registrations \
  -H "Content-Type: application/json" \
  -d '{
    "studentId": 1,
    "examId": 1,
    "courseId": 1
  }'

# 4. Check registration status
curl http://localhost:8086/api/exam-registrations/1
```

## Database Schema

Each service has its own database schema that is automatically migrated on startup via GORM AutoMigrate.

**Databases:**
- `user_service` - User profiles
- `exam_service` - Exams, examination periods, exam students, exam results
- `class_service` - Students, departments, study programs, courses
- `transaction_service` - Wallets, transactions
- `exam_registration_service` - Exam registrations, saga items

## Development

### Adding a New Service

1. Create service directory: `my-service-go/`
2. Initialize Go module: `go mod init my-service`
3. Add dependencies:
   ```bash
   go get github.com/gin-gonic/gin
   go get gorm.io/gorm
   go get gorm.io/driver/mysql
   go get github.com/rabbitmq/amqp091-go
   ```
4. Implement clean architecture structure
5. Add Dockerfile
6. Update `docker-compose-go.yml`

### Running Tests

```bash
cd service-name-go
go test ./...
```

### Building Binaries

```bash
cd service-name-go
go build -o bin/service cmd/main.go
```

## Monitoring & Debugging

### RabbitMQ Management
- URL: http://localhost:15672
- Credentials: guest/guest
- Monitor queues, exchanges, and message flow

### Mailhog
- URL: http://localhost:8025
- View all emails sent by the system

### Service Health Checks
Each service exposes a `/health` endpoint:
```bash
curl http://localhost:8080/health
```

### Logs
View service logs:
```bash
docker-compose -f docker-compose-go.yml logs -f user-service-go
docker-compose -f docker-compose-go.yml logs -f exam-registration-service-go
```

## Configuration

Each service can be configured via environment variables or `.env` files:

**Common Variables:**
- `DB_HOST` - MySQL host
- `DB_PORT` - MySQL port
- `DB_USER` - MySQL user
- `DB_PASSWORD` - MySQL password
- `DB_NAME` - Database name
- `SERVER_PORT` - Service port
- `RABBITMQ_URL` - RabbitMQ connection URL

## Stopping Services

```bash
docker-compose -f docker-compose-go.yml down

# Remove volumes (database data)
docker-compose -f docker-compose-go.yml down -v
```

## Advantages of Go Rewrite

1. **Performance:** Go's compiled nature provides faster execution
2. **Concurrency:** Native goroutines for handling concurrent operations
3. **Memory Efficiency:** Lower memory footprint compared to PHP
4. **Type Safety:** Strong static typing catches errors at compile time
5. **Simple Deployment:** Single binary deployment per service
6. **Better Tooling:** Built-in testing, profiling, and race detection

## Migration Notes

This Go version maintains **feature parity** with the original PHP/Symfony implementation:
- ✅ All entities and relationships preserved
- ✅ Saga pattern implementation maintained
- ✅ Message-driven architecture intact
- ✅ API contracts compatible
- ✅ Same port assignments for easy transition

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Write tests
5. Submit a pull request

## License

[Add your license here]

## Support

For issues and questions, please open an issue on the GitHub repository.
