package repository

import (
	"transaction-service/internal/domain"

	"gorm.io/gorm"
)

type WalletRepository interface {
	Create(wallet *domain.Wallet) error
	FindByID(id uint) (*domain.Wallet, error)
	FindByStudentID(studentID int) (*domain.Wallet, error)
	FindAll(offset, limit int) ([]domain.Wallet, error)
	Update(wallet *domain.Wallet) error
	Delete(id uint) error
	Count() (int64, error)
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{db: db}
}

func (r *walletRepository) Create(wallet *domain.Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *walletRepository) FindByID(id uint) (*domain.Wallet, error) {
	var wallet domain.Wallet
	err := r.db.Preload("Transactions").First(&wallet, id).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *walletRepository) FindByStudentID(studentID int) (*domain.Wallet, error) {
	var wallet domain.Wallet
	err := r.db.Preload("Transactions").Where("student_id = ?", studentID).First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *walletRepository) FindAll(offset, limit int) ([]domain.Wallet, error) {
	var wallets []domain.Wallet
	err := r.db.Offset(offset).Limit(limit).Find(&wallets).Error
	return wallets, err
}

func (r *walletRepository) Update(wallet *domain.Wallet) error {
	return r.db.Save(wallet).Error
}

func (r *walletRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Wallet{}, id).Error
}

func (r *walletRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&domain.Wallet{}).Count(&count).Error
	return count, err
}
