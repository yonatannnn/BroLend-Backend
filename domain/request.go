// domain/request.go
package domain

type RequestStatus string

const (
	RequestPending  RequestStatus = "pending"
	RequestAccepted RequestStatus = "accepted"
	RequestRejected RequestStatus = "rejected"
)

type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyETB Currency = "ETB" // Birr
)

type Request struct {
	ID         string        `bson:"_id" json:"id"`
	RequestID  string        `bson:"requestId" json:"requestId"`
	SenderID   string        `bson:"senderId" json:"senderId"`
	ReceiverID string        `bson:"receiverId" json:"receiverId"`
	Status     RequestStatus `bson:"status" json:"status"`
	Amount     float64       `bson:"amount" json:"amount"`
	Currency   Currency      `bson:"currency" json:"currency"`
	Date       int64         `bson:"date" json:"date"` // Unix timestamp
}
