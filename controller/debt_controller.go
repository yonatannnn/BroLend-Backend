package controller

import (
	"brolend/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DebtController struct {
	du domain.DebtUsecase
}

func NewDebtController(du domain.DebtUsecase) *DebtController {
	return &DebtController{du: du}
}

// POST /debt
func (dc *DebtController) CreateDebt(c *gin.Context) {
	var debt domain.Debt
	if err := c.ShouldBindJSON(&debt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Get borrower ID from JWT claims
	borrowerID := c.GetString("user_id")
	debt.BorrowerID = borrowerID

	id, err := dc.du.CreateDebt(debt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Debt created", "id": id})
}

// POST /debt/:id/accept
func (dc *DebtController) AcceptDebt(c *gin.Context) {
	debtID := c.Param("id")
	lenderID := c.GetString("user_id") // from JWT claims
	if err := dc.du.AcceptDebt(debtID, lenderID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Debt accepted"})
}

// POST /debt/:id/reject
func (dc *DebtController) RejectDebt(c *gin.Context) {
	debtID := c.Param("id")
	lenderID := c.GetString("user_id")
	if err := dc.du.RejectDebt(debtID, lenderID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Debt rejected"})
}

// POST /debt/:id/request-paid
func (dc *DebtController) RequestPaidApproval(c *gin.Context) {
	debtID := c.Param("id")
	borrowerID := c.GetString("user_id")
	if err := dc.du.RequestPaidApproval(debtID, borrowerID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Paid approval requested"})
}

// POST /debt/:id/approve-payment
func (dc *DebtController) ApprovePayment(c *gin.Context) {
	debtID := c.Param("id")
	lenderID := c.GetString("user_id")
	if err := dc.du.ApprovePayment(debtID, lenderID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Payment approved"})
}

// POST /debt/:id/reject-payment
func (dc *DebtController) RejectPaymentRequest(c *gin.Context) {
	debtID := c.Param("id")
	lenderID := c.GetString("user_id")
	if err := dc.du.RejectPaymentRequest(debtID, lenderID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Payment request rejected"})
}

// GET /debt/net
func (dc *DebtController) GetNetAmounts(c *gin.Context) {
	userID := c.GetString("user_id")
	net, err := dc.du.GetNetAmounts(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, net)
}

// GET /debt/history
func (dc *DebtController) GetHistory(c *gin.Context) {
	userID := c.GetString("user_id")
	history, err := dc.du.GetHistory(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, history)
}

// GET /debt/active-incoming
func (dc *DebtController) GetActiveIncoming(c *gin.Context) {
	lenderID := c.GetString("user_id")
	debts, err := dc.du.GetActiveIncoming(lenderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, debts)
}

// GET /debt/active-outgoing
func (dc *DebtController) GetActiveOutgoing(c *gin.Context) {
	borrowerID := c.GetString("user_id")
	debts, err := dc.du.GetActiveOutgoing(borrowerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, debts)
}

// GET /debt/incoming-requests
func (dc *DebtController) GetIncomingRequests(c *gin.Context) {
	lenderID := c.GetString("user_id")
	debts, err := dc.du.GetIncomingRequests(lenderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, debts)
}
