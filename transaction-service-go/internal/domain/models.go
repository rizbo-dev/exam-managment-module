package domain

import "time"

type Wallet struct {
	ID           uint          `json:"id" gorm:"primaryKey"`
	StudentID    int           `json:"studentId" gorm:"column:student_id;not null;unique"`
	Balance      float64       `json:"balance" gorm:"not null;default:0"`
	Transactions []Transaction `json:"transactions,omitempty" gorm:"foreignKey:WalletID"`
	CreatedAt    time.Time     `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time     `json:"updatedAt" gorm:"autoUpdateTime"`
}

type Transaction struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	WalletID  uint      `json:"walletId" gorm:"column:wallet_id;not null"`
	Wallet    *Wallet   `json:"wallet,omitempty" gorm:"foreignKey:WalletID"`
	Type      string    `json:"type" gorm:"size:50;not null"` // deposit, withdrawal, exam_fee
	Amount    float64   `json:"amount" gorm:"not null"`
	Status    string    `json:"status" gorm:"size:50;not null;default:completed"` // completed, pending, failed
	Reference string    `json:"reference" gorm:"size:255"` // Reference ID for saga or external system
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// Request DTOs
type CreateWalletRequest struct {
	StudentID int     `json:"studentId" binding:"required"`
	Balance   float64 `json:"balance"`
}

type DepositRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

type WithdrawRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}
