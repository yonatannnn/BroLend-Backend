package usecase

import "brolend/domain"

type DebtUsecase struct {
    repo domain.DebtRepository
}

func NewDebtUsecase(repo domain.DebtRepository) *DebtUsecase {
    return &DebtUsecase{repo: repo}
}

// Implement domain.DebtUsecase interface

func (u *DebtUsecase) CreateDebt(debt domain.Debt) (string, error) {
    // TODO: business logic for creating a debt
    return "", nil
}

func (u *DebtUsecase) AcceptDebt(debtID string, lenderID string) error {
    // TODO: business logic for accepting a debt
    return nil
}

func (u *DebtUsecase) RejectDebt(debtID string, lenderID string) error {
    // TODO: business logic for rejecting a debt
    return nil
}

func (u *DebtUsecase) RequestPaidApproval(debtID string, borrowerID string) error {
    // TODO: business logic for borrower requesting paid approval
    return nil
}

func (u *DebtUsecase) ApprovePayment(debtID string, lenderID string) error {
    // TODO: business logic for lender approving payment
    return nil
}

func (u *DebtUsecase) RejectPaymentRequest(debtID string, lenderID string) error {
    // TODO: business logic for lender rejecting payment request
    return nil
}

func (u *DebtUsecase) GetNetAmounts(userID string) (map[domain.Currency]float64, error) {
    // TODO: business logic for net amounts
    return nil, nil
}

func (u *DebtUsecase) GetHistory(userID string) ([]*domain.Debt, error) {
    // TODO: business logic for user debt history
    return nil, nil
}

func (u *DebtUsecase) GetActiveIncoming(lenderID string) ([]*domain.Debt, error) {
    // TODO: business logic for active incoming debts
    return nil, nil
}

func (u *DebtUsecase) GetActiveOutgoing(borrowerID string) ([]*domain.Debt, error) {
    // TODO: business logic for active outgoing debts
    return nil, nil
}

func (u *DebtUsecase) GetIncomingRequests(lenderID string) ([]*domain.Debt, error) {
    // TODO: business logic for incoming requests
    return nil, nil
}