package messaging

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewConsumer(rabbitmqURL string) (*Consumer, error) {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn:    conn,
		channel: channel,
	}, nil
}

func (c *Consumer) StartConsuming() error {
	// Setup student update queue
	err := c.channel.ExchangeDeclare("student-update", "fanout", true, false, false, false, nil)
	if err != nil {
		return err
	}

	queue, err := c.channel.QueueDeclare("student-update", true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = c.channel.QueueBind(queue.Name, "", "student-update", false, nil)
	if err != nil {
		return err
	}

	msgs, err := c.channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			c.handleStudentUpdate(msg)
		}
	}()

	log.Println("Started consuming email service messages")
	return nil
}

func (c *Consumer) handleStudentUpdate(msg amqp.Delivery) {
	var message struct {
		StudentID  int    `json:"studentId"`
		FirstName  string `json:"firstName"`
		LastName   string `json:"lastName"`
		Email      string `json:"email"`
	}

	if err := json.Unmarshal(msg.Body, &message); err != nil {
		log.Printf("Failed to unmarshal message: %v", err)
		msg.Nack(false, false)
		return
	}

	log.Printf("Sending email notification to student %d (%s %s) at %s",
		message.StudentID, message.FirstName, message.LastName, message.Email)

	// Here you would integrate with an email service (e.g., SMTP, SendGrid, etc.)
	// For now, we'll just log it
	log.Printf("Email sent successfully to %s", message.Email)

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
