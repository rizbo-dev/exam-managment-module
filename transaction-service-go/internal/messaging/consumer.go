package messaging

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"transaction-service/internal/domain"
	"transaction-service/internal/repository"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn            *amqp.Connection
	channel         *amqp.Channel
	walletRepo      repository.WalletRepository
	transactionRepo repository.TransactionRepository
	publisher       *Publisher
}

func NewConsumer(
	rabbitmqURL string,
	walletRepo repository.WalletRepository,
	transactionRepo repository.TransactionRepository,
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
		conn:            conn,
		channel:         channel,
		walletRepo:      walletRepo,
		transactionRepo: transactionRepo,
		publisher:       publisher,
	}, nil
}

func (c *Consumer) StartConsuming() error {
	// Setup wallet validation queue
	if err := c.setupWalletValidationQueue(); err != nil {
		return err
	}

	// Setup wallet insert queue
	if err := c.setupWalletInsertQueue(); err != nil {
		return err
	}

	log.Println("Started consuming transaction service messages")
	return nil
}

func (c *Consumer) setupWalletValidationQueue() error {
	err := c.channel.ExchangeDeclare("user-wallet-validation", "fanout", true, false, false, false, nil)
	if err != nil {
		return err
	}

	queue, err := c.channel.QueueDeclare("user-wallet-validation", true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = c.channel.QueueBind(queue.Name, "", "user-wallet-validation", false, nil)
	if err != nil {
		return err
	}

	msgs, err := c.channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			c.handleWalletValidation(msg)
		}
	}()

	return nil
}

func (c *Consumer) setupWalletInsertQueue() error {
	err := c.channel.ExchangeDeclare("user-wallet-insert", "fanout", true, false, false, false, nil)
	if err != nil {
		return err
	}

	queue, err := c.channel.QueueDeclare("user-wallet-insert", true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = c.channel.QueueBind(queue.Name, "", "user-wallet-insert", false, nil)
	if err != nil {
		return err
	}

	msgs, err := c.channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			c.handleWalletInsert(msg)
		}
	}()

	return nil
}

func (c *Consumer) handleWalletValidation(msg amqp.Delivery) {
	var message struct {
		StudentID int  `json:"studentId"`
		ExamID    uint `json:"examId"`
		SagaItemID uint `json:"sagaItemId"`
	}

	if err := json.Unmarshal(msg.Body, &message); err != nil {
		log.Printf("Failed to unmarshal message: %v", err)
		msg.Nack(false, false)
		return
	}

	log.Printf("Validating wallet for student %d and exam %d", message.StudentID, message.ExamID)

	// Get wallet
	wallet, err := c.walletRepo.FindByStudentID(message.StudentID)
	if err != nil {
		log.Printf("Wallet not found: %v", err)
		c.publisher.PublishSagaResponse(message.SagaItemID, false, "Wallet not found")
		msg.Ack(false)
		return
	}

	// Get exam cost from exam service
	examCost, err := c.getExamCost(message.ExamID)
	if err != nil {
		log.Printf("Failed to get exam cost: %v", err)
		c.publisher.PublishSagaResponse(message.SagaItemID, false, "Failed to get exam cost")
		msg.Ack(false)
		return
	}

	// Validate balance
	if wallet.Balance < examCost {
		log.Printf("Insufficient balance: have %.2f, need %.2f", wallet.Balance, examCost)
		c.publisher.PublishSagaResponse(message.SagaItemID, false, fmt.Sprintf("Insufficient balance"))
		msg.Ack(false)
		return
	}

	log.Printf("Wallet validation successful")
	c.publisher.PublishSagaResponse(message.SagaItemID, true, "Wallet has sufficient balance")
	msg.Ack(false)
}

func (c *Consumer) handleWalletInsert(msg amqp.Delivery) {
	var message struct {
		StudentID int  `json:"studentId"`
		ExamID    uint `json:"examId"`
		SagaItemID uint `json:"sagaItemId"`
	}

	if err := json.Unmarshal(msg.Body, &message); err != nil {
		log.Printf("Failed to unmarshal message: %v", err)
		msg.Nack(false, false)
		return
	}

	log.Printf("Deducting exam fee for student %d and exam %d", message.StudentID, message.ExamID)

	// Get wallet
	wallet, err := c.walletRepo.FindByStudentID(message.StudentID)
	if err != nil {
		log.Printf("Wallet not found: %v", err)
		c.publisher.PublishSagaResponse(message.SagaItemID, false, "Wallet not found")
		msg.Ack(false)
		return
	}

	// Get exam cost
	examCost, err := c.getExamCost(message.ExamID)
	if err != nil {
		log.Printf("Failed to get exam cost: %v", err)
		c.publisher.PublishSagaResponse(message.SagaItemID, false, "Failed to get exam cost")
		msg.Ack(false)
		return
	}

	// Check balance again
	if wallet.Balance < examCost {
		log.Printf("Insufficient balance")
		c.publisher.PublishSagaResponse(message.SagaItemID, false, "Insufficient balance")
		msg.Ack(false)
		return
	}

	// Deduct balance
	wallet.Balance -= examCost
	if err := c.walletRepo.Update(wallet); err != nil {
		log.Printf("Failed to update wallet: %v", err)
		c.publisher.PublishSagaResponse(message.SagaItemID, false, "Failed to deduct balance")
		msg.Ack(false)
		return
	}

	// Create transaction record
	transaction := &domain.Transaction{
		WalletID:  wallet.ID,
		Type:      "exam_fee",
		Amount:    -examCost,
		Status:    "completed",
		Reference: fmt.Sprintf("exam_%d_saga_%d", message.ExamID, message.SagaItemID),
	}

	if err := c.transactionRepo.Create(transaction); err != nil {
		log.Printf("Failed to create transaction: %v", err)
	}

	log.Printf("Successfully deducted exam fee: %.2f", examCost)
	c.publisher.PublishSagaResponse(message.SagaItemID, true, fmt.Sprintf("Deducted %.2f from wallet", examCost))
	msg.Ack(false)
}

func (c *Consumer) getExamCost(examID uint) (float64, error) {
	// Call exam service to get exam cost
	url := fmt.Sprintf("http://exam-nginx-service/api/exams/%d", examID)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("exam service returned status %d", resp.StatusCode)
	}

	var examResp struct {
		Cost float64 `json:"cost"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&examResp); err != nil {
		return 0, err
	}

	return examResp.Cost, nil
}

func (c *Consumer) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
}
