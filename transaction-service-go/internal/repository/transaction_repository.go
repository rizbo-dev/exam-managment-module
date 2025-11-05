package repository

import (
	"transaction-service/internal/domain"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *domain.Transaction) error
	FindByID(id uint) (*domain.Transaction, error)
	FindByWalletID(walletID uint) ([]domain.Transaction, error)
	FindAll(offset, limit int) ([]domain.Transaction, error)
	Count() (int64, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(transaction *domain.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) FindByID(id uint) (*domain.Transaction, error) {
	var transaction domain.Transaction
	err := r.db.Preload("Wallet").First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *transactionRepository) FindByWalletID(walletID uint) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Where("wallet_id = ?", walletID).Order("created_at DESC").Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) FindAll(offset, limit int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Preload("Wallet").Offset(offset).Limit(limit).Order("created_at DESC").Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&domain.Transaction{}).Count(&count).Error
	return count, err
}
