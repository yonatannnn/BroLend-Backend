// domain/repositories.go
package domain

type UserRepository interface {
	Create(user *User) error
	FindByID(id string) (*User, error)
	FindByUsername(username string) (*User, error)
	FindAllUsers() ([]*User, error)
}

type RequestRepository interface {
	Create(request *Request) error
	FindByID(id string) (*Request, error)
	FindByReceiverID(receiverID string, status RequestStatus) ([]*Request, error)
	UpdateStatus(id string, status RequestStatus) error
}

type DebtRepository interface {
	Create(debt *Debt) error
	GetDebtSummary(user1ID, user2ID string) (*DebtSummary, error)
	FindByDeptID(debtID string) (*Debt, error)
	UpdatePayment(debtID string, amount float64) error
	FindUnsettledByUserPair(user1ID, user2ID, currency string) ([]*Debt, error)
}

type PaymentRepository interface {
	Create(payment *Payment) error
	FindByID(id string) (*Payment, error)
}
