package domain

type UserUsecase interface {
	Create(user *User) (string, error)
	Login(username, password string) (string, *User, error)
}

type RequestUsecase interface {
	Create(string, string, float64, Currency) (*Request, error)
	FindByID(id string) (*Request, error)
	FindByReceiverID(receiverID string, status RequestStatus) ([]*Request, error)
	Accept(id string) error
	Reject(id string) error
}

type DebtUsecase interface {
	Create(debt *Debt) error
	FindBetweenUsers(user1ID, user2ID string) (*DebtSummary, error)
	UpdatePayment(amount float64, borrowerID string, lenderID string) error
	FindUnsettledByUserPair(user1ID, user2ID, currency string) ([]*Debt, error)
}

type PaymentUsecase interface {
	Create(payment *Payment) error
}
