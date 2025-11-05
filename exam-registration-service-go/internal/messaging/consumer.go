package messaging

import (
	"encoding/json"
	"exam-registration-service/internal/service"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn        *amqp.Connection
	channel     *amqp.Channel
	orchestrator *service.SagaOrchestrator
}

func NewConsumer(rabbitmqURL string, orchestrator *service.SagaOrchestrator) (*Consumer, error) {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn:        conn,
		channel:     channel,
		orchestrator: orchestrator,
	}, nil
}

func (c *Consumer) StartConsuming() error {
	// Setup response saga item queue
	err := c.channel.ExchangeDeclare("response-saga-item", "fanout", true, false, false, false, nil)
	if err != nil {
		return err
	}

	queue, err := c.channel.QueueDeclare("response-saga-item", true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = c.channel.QueueBind(queue.Name, "", "response-saga-item", false, nil)
	if err != nil {
		return err
	}

	msgs, err := c.channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			c.handleSagaResponse(msg)
		}
	}()

	log.Println("Started consuming saga response messages")
	return nil
}

func (c *Consumer) handleSagaResponse(msg amqp.Delivery) {
	var response struct {
		SagaItemID uint   `json:"sagaItemId"`
		Success    bool   `json:"success"`
		Message    string `json:"message"`
	}

	if err := json.Unmarshal(msg.Body, &response); err != nil {
		log.Printf("Failed to unmarshal saga response: %v", err)
		msg.Nack(false, false)
		return
	}

	log.Printf("Received saga response for item %d: success=%v, message=%s",
		response.SagaItemID, response.Success, response.Message)

	if err := c.orchestrator.HandleSagaResponse(response.SagaItemID, response.Success, response.Message); err != nil {
		log.Printf("Failed to handle saga response: %v", err)
	}

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
