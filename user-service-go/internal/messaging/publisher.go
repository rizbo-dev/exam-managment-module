package messaging

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"user-service/internal/domain"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewPublisher(rabbitmqURL string) (*Publisher, error) {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Publisher{
		conn:    conn,
		channel: channel,
	}, nil
}

func (p *Publisher) PublishUserUpdated(user *domain.User) error {
	message := map[string]interface{}{
		"userId":         user.ID,
		"firstName":      user.FirstName,
		"lastName":       user.LastName,
		"role":           user.Role,
		"departmentId":   user.DepartmentID,
		"studyProgramId": user.StudyProgramID,
	}

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = p.channel.PublishWithContext(ctx,
		"user-updated", // exchange
		"",             // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		log.Printf("Failed to publish user updated event: %v", err)
		return err
	}

	log.Printf("Published user updated event for user ID: %d", user.ID)
	return nil
}

func (p *Publisher) Close() {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.conn != nil {
		p.conn.Close()
	}
}
