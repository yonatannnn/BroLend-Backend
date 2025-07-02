package usecase

import (
    "brolend/domain"
    "errors"
    "time"
)

type DebtUsecase struct {
    repo domain.DebtRepository
}

func NewDebtUsecase(repo domain.DebtRepository) *DebtUsecase {
    return &DebtUsecase{repo: repo}
}

func (u *DebtUsecase) CreateDebt(debt domain.Debt) (string, error) {
    // Set initial status and date
    debt.Status = domain.RequestPending
    debt.Date = time.Now().Unix()
    return u.repo.Create(&debt)
}

func (u *DebtUsecase) AcceptDebt(debtID string, lenderID string) error {
    debt, err := u.repo.FindByID(debtID)
    if err != nil {
        return err
    }
    if debt.LenderID != lenderID {
        return errors.New("not authorized")
    }
    if debt.Status != domain.RequestPending {
        return errors.New("debt is not pending")
    }
    return u.repo.UpdateStatus(debtID, domain.RequestAccepted)
}

func (u *DebtUsecase) RejectDebt(debtID string, lenderID string) error {
    debt, err := u.repo.FindByID(debtID)
    if err != nil {
        return err
    }
    if debt.LenderID != lenderID {
        return errors.New("not authorized")
    }
    if debt.Status != domain.RequestPending {
        return errors.New("debt is not pending")
    }
    return u.repo.UpdateStatus(debtID, domain.RequestRejected)
}

func (u *DebtUsecase) RequestPaidApproval(debtID string, borrowerID string) error {
    debt, err := u.repo.FindByID(debtID)
    if err != nil {
        return err
    }
    if debt.BorrowerID != borrowerID {
        return errors.New("not authorized")
    }
    if debt.Status != domain.RequestAccepted {
        return errors.New("debt is not accepted")
    }
    return u.repo.UpdateStatus(debtID, domain.RequestPaid)
}

func (u *DebtUsecase) ApprovePayment(debtID string, lenderID string) error {
    debt, err := u.repo.FindByID(debtID)
    if err != nil {
        return err
    }
    if debt.LenderID != lenderID {
        return errors.New("not authorized")
    }
    if debt.Status != domain.RequestPaid {
        return errors.New("payment not requested")
    }
    return u.repo.UpdateStatus(debtID, domain.Settled)
}

func (u *DebtUsecase) RejectPaymentRequest(debtID string, lenderID string) error {
    debt, err := u.repo.FindByID(debtID)
    if err != nil {
        return err
    }
    if debt.LenderID != lenderID {
        return errors.New("not authorized")
    }
    if debt.Status != domain.RequestPaid {
        return errors.New("payment not requested")
    }
    return u.repo.UpdateStatus(debtID, domain.RequestAccepted) // revert to accepted
}

func (u *DebtUsecase) GetNetAmounts(userID string) (map[domain.Currency]float64, error) {
    return u.repo.FindNetAmounts(userID)
}

func (u *DebtUsecase) GetHistory(userID string) ([]*domain.Debt, error) {
    return u.repo.FindHistory(userID)
}

func (u *DebtUsecase) GetActiveIncoming(lenderID string) ([]*domain.Debt, error) {
    return u.repo.FindActiveIncoming(lenderID)
}

func (u *DebtUsecase) GetActiveOutgoing(borrowerID string) ([]*domain.Debt, error) {
    return u.repo.FindActiveOutgoing(borrowerID)
}

func (u *DebtUsecase) GetIncomingRequests(lenderID string) ([]*domain.Debt, error) {
    return u.repo.FindIncomingRequests(lenderID)
}