package handler

import (
	"net/http"
	"strconv"
	"transaction-service/internal/domain"
	"transaction-service/internal/repository"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	walletRepo      repository.WalletRepository
	transactionRepo repository.TransactionRepository
}

func NewTransactionHandler(
	walletRepo repository.WalletRepository,
	transactionRepo repository.TransactionRepository,
) *TransactionHandler {
	return &TransactionHandler{
		walletRepo:      walletRepo,
		transactionRepo: transactionRepo,
	}
}

// Wallet endpoints
func (h *TransactionHandler) CreateWallet(c *gin.Context) {
	var req domain.CreateWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wallet := &domain.Wallet{
		StudentID: req.StudentID,
		Balance:   req.Balance,
	}

	if err := h.walletRepo.Create(wallet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, wallet)
}

func (h *TransactionHandler) GetWalletByStudentID(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("studentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	wallet, err := h.walletRepo.FindByStudentID(studentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	c.JSON(http.StatusOK, wallet)
}

func (h *TransactionHandler) GetWallets(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "30"))

	offset := (page - 1) * limit

	wallets, err := h.walletRepo.FindAll(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, _ := h.walletRepo.Count()

	c.JSON(http.StatusOK, gin.H{
		"data":       wallets,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}

func (h *TransactionHandler) Deposit(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("studentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var req domain.DepositRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wallet, err := h.walletRepo.FindByStudentID(studentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	wallet.Balance += req.Amount
	if err := h.walletRepo.Update(wallet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create transaction record
	transaction := &domain.Transaction{
		WalletID: wallet.ID,
		Type:     "deposit",
		Amount:   req.Amount,
		Status:   "completed",
	}
	h.transactionRepo.Create(transaction)

	c.JSON(http.StatusOK, wallet)
}

// Transaction endpoints
func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "30"))

	offset := (page - 1) * limit

	transactions, err := h.transactionRepo.FindAll(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, _ := h.transactionRepo.Count()

	c.JSON(http.StatusOK, gin.H{
		"data":       transactions,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}

func (h *TransactionHandler) GetWalletTransactions(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("studentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	wallet, err := h.walletRepo.FindByStudentID(studentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	transactions, err := h.transactionRepo.FindByWalletID(wallet.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
