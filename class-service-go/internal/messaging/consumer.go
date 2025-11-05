package messaging

import (
	"class-service/internal/repository"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn                   *amqp.Connection
	channel                *amqp.Channel
	studentRepo            repository.StudentRepository
	studyProgramRepo       repository.StudyProgramRepository
	courseRepo             repository.CourseRepository
	publisher              *Publisher
}

func NewConsumer(
	rabbitmqURL string,
	studentRepo repository.StudentRepository,
	studyProgramRepo repository.StudyProgramRepository,
	courseRepo repository.CourseRepository,
	publisher *Publisher,
) (*Consumer, error) {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn:             conn,
		channel:          channel,
		studentRepo:      studentRepo,
		studyProgramRepo: studyProgramRepo,
		courseRepo:       courseRepo,
		publisher:        publisher,
	}, nil
}

func (c *Consumer) StartConsuming() error {
	// Setup user class verification queue
	err := c.setupUserClassVerificationQueue()
	if err != nil {
		return err
	}

	log.Println("Started consuming class service messages")
	return nil
}

func (c *Consumer) setupUserClassVerificationQueue() error {
	// Declare exchange
	err := c.channel.ExchangeDeclare(
		"user-class-verification",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Declare queue
	queue, err := c.channel.QueueDeclare(
		"user-class-verification",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Bind queue to exchange
	err = c.channel.QueueBind(
		queue.Name,
		"",
		"user-class-verification",
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Start consuming
	msgs, err := c.channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			c.handleUserClassVerification(msg)
		}
	}()

	return nil
}

func (c *Consumer) handleUserClassVerification(msg amqp.Delivery) {
	var message struct {
		StudentID      int  `json:"studentId"`
		CourseID       int  `json:"courseId"`
		SagaItemID     uint `json:"sagaItemId"`
	}

	if err := json.Unmarshal(msg.Body, &message); err != nil {
		log.Printf("Failed to unmarshal message: %v", err)
		msg.Nack(false, false)
		return
	}

	log.Printf("Verifying student %d enrollment for course %d", message.StudentID, message.CourseID)

	// Get student
	student, err := c.studentRepo.FindByID(uint(message.StudentID))
	if err != nil {
		log.Printf("Student not found: %v", err)
		c.publisher.PublishSagaResponse(message.SagaItemID, false, "Student not found")
		msg.Ack(false)
		return
	}

	// Get courses for student's study program
	courses, err := c.courseRepo.FindByStudyProgramID(student.StudyProgramID)
	if err != nil {
		log.Printf("Failed to get courses: %v", err)
		c.publisher.PublishSagaResponse(message.SagaItemID, false, "Failed to verify enrollment")
		msg.Ack(false)
		return
	}

	// Check if course is in student's study program
	enrolled := false
	for _, course := range courses {
		if int(course.ID) == message.CourseID {
			enrolled = true
			break
		}
	}

	if !enrolled {
		log.Printf("Student not enrolled in course")
		c.publisher.PublishSagaResponse(message.SagaItemID, false, "Student not enrolled in this course")
		msg.Ack(false)
		return
	}

	log.Printf("Student successfully verified for course")
	c.publisher.PublishSagaResponse(message.SagaItemID, true, "Student enrolled in course")
	msg.Ack(false)
}

func (c *Consumer) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
}
