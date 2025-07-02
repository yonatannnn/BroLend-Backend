// domain/debt.go
package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Debt struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	DebtID     string             `bson:"debt_id" json:"debt_id"`
	LenderID   string             `bson:"lender_id" json:"lender_id"`
	BorrowerID string             `bson:"borrower_id" json:"borrower_id"`
	Amount     float64            `bson:"amount" json:"amount"`
	Currency   Currency           `bson:"currency" json:"currency"`
	Date       int64              `bson:"date" json:"date"`
	Status     Status             `bson:"status" json:"status"`
}

type DebtRepository interface {
	Create(debt *Debt) (string, error)
	FindByID(id string) (*Debt, error)
	UpdateStatus(id string, status Status) error
	Update(debt *Debt) error
	FindByLenderID(lenderID string) ([]*Debt, error)
	FindByBorrowerID(borrowerID string) ([]*Debt, error)
	FindActiveIncoming(lenderID string) ([]*Debt, error)
	FindActiveOutgoing(borrowerID string) ([]*Debt, error)
	FindHistory(userID string) ([]*Debt, error)
	FindNetAmounts(userID string) (map[Currency]float64, error)
	FindIncomingRequests(lenderID string) ([]*Debt, error)
}

type DebtUsecase interface {
	CreateDebt(debt Debt) (string, error)
	AcceptDebt(debtID string, lenderID string) error
	RejectDebt(debtID string, lenderID string) error
	RequestPaidApproval(debtID string, borrowerID string) error
	ApprovePayment(debtID string, lenderID string) error
	RejectPaymentRequest(debtID string, lenderID string) error
	GetNetAmounts(userID string) (map[Currency]float64, error)
	GetHistory(userID string) ([]*Debt, error)
	GetActiveIncoming(lenderID string) ([]*Debt, error)
	GetActiveOutgoing(borrowerID string) ([]*Debt, error)
	GetIncomingRequests(lenderID string) ([]*Debt, error)
}

type Status string

const (
	RequestPending  Status = "R"
	RequestAccepted Status = "A"
	RequestRejected Status = "RJ"
	RequestPaid     Status = "RP"
	Settled         Status = "S"
)

type Currency string

const (
	CurrencyUSD  Currency = "USD"
	CurrencyETB  Currency = "ETB" // Birr
	CurrencyUSDT Currency = "USDT"
)
