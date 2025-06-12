// domain/debt.go
package domain

type Debt struct {
	ID         string   `bson:"_id" json:"id"`
	Lender     string   `bson:"lender" json:"lender"`
	Borrower   string   `bson:"borrower" json:"borrower"`
	Amount     float64  `bson:"amount" json:"amount"`
	Currency   Currency `bson:"currency" json:"currency"`
	Date       int64    `bson:"date" json:"date"`
	PaidAmount float64  `bson:"paidAmount" json:"paidAmount"`
	Remaining  float64  `bson:"remaining" json:"remaining"`
	IsSettled  bool     `bson:"isSettled" json:"isSettled"`
	RequestID  string   `bson:"requestId" json:"requestId"`
}

type DebtSummary struct {
	SettledDebts        []Debt
	UnsettledAsLender   []Debt
	UnsettledAsBorrower []Debt
	NetETB              float64
	NetUSD              float64
}
