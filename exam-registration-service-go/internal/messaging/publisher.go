package messaging

import (
	"context"
	"encoding/json"
	"log"
	"time"

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

func (p *Publisher) PublishUserClassVerification(sagaItemID uint, studentID int, courseID int) error {
	message := map[string]interface{}{
		"sagaItemId": sagaItemID,
		"studentId":  studentID,
		"courseId":   courseID,
	}
	return p.publishToExchange("user-class-verification", message)
}

func (p *Publisher) PublishUserWalletValidation(sagaItemID uint, studentID int, examID uint) error {
	message := map[string]interface{}{
		"sagaItemId": sagaItemID,
		"studentId":  studentID,
		"examId":     examID,
	}
	return p.publishToExchange("user-wallet-validation", message)
}

func (p *Publisher) PublishUserWalletInsert(sagaItemID uint, studentID int, examID uint) error {
	message := map[string]interface{}{
		"sagaItemId": sagaItemID,
		"studentId":  studentID,
		"examId":     examID,
	}
	return p.publishToExchange("user-wallet-insert", message)
}

func (p *Publisher) PublishExamRegistration(sagaItemID uint, studentID int, examID uint, examRegistrationID uint) error {
	message := map[string]interface{}{
		"sagaItemId":         sagaItemID,
		"studentId":          studentID,
		"examId":             examID,
		"examRegistrationId": examRegistrationID,
	}
	return p.publishToExchange("exam-registration", message)
}

func (p *Publisher) publishToExchange(exchange string, message map[string]interface{}) error {
	// Declare exchange
	err := p.channel.ExchangeDeclare(exchange, "fanout", true, false, false, false, nil)
	if err != nil {
		return err
	}

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = p.channel.PublishWithContext(ctx, exchange, "", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})

	if err != nil {
		log.Printf("Failed to publish to %s: %v", exchange, err)
		return err
	}

	log.Printf("Published message to %s", exchange)
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
