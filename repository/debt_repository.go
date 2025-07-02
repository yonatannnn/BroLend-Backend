package repository

import (
	"brolend/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DebtRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewDebtRepository(col *mongo.Collection, ctx context.Context) *DebtRepository {
	return &DebtRepository{collection: col, ctx: ctx}
}

func (r *DebtRepository) Create(debt *domain.Debt) (string, error) {
	debt.ID = primitive.NewObjectID()
	debt.DebtID = debt.ID.Hex()
	res, err := r.collection.InsertOne(r.ctx, debt)
	if err != nil {
		return "", err
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if ok {
		debt.ID = oid
		return oid.Hex(), nil
	}
	return "", nil
}

func (r *DebtRepository) FindByID(id string) (*domain.Debt, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var debt domain.Debt
	err = r.collection.FindOne(r.ctx, bson.M{"_id": oid}).Decode(&debt)
	if err != nil {
		return nil, err
	}
	return &debt, nil
}

func (r *DebtRepository) UpdateStatus(id string, status domain.Status) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.UpdateOne(r.ctx, bson.M{"_id": oid}, bson.M{"$set": bson.M{"status": status}})
	return err
}

func (r *DebtRepository) Update(debt *domain.Debt) error {
	_, err := r.collection.ReplaceOne(r.ctx, bson.M{"_id": debt.ID}, debt)
	return err
}

func (r *DebtRepository) FindByLenderID(lenderID string) ([]*domain.Debt, error) {
	cursor, err := r.collection.Find(r.ctx, bson.M{"lender_id": lenderID})
	if err != nil {
		return nil, err
	}
	var debts []*domain.Debt
	if err = cursor.All(r.ctx, &debts); err != nil {
		return nil, err
	}
	return debts, nil
}

func (r *DebtRepository) FindByBorrowerID(borrowerID string) ([]*domain.Debt, error) {
	cursor, err := r.collection.Find(r.ctx, bson.M{"borrower_id": borrowerID})
	if err != nil {
		return nil, err
	}
	var debts []*domain.Debt
	if err = cursor.All(r.ctx, &debts); err != nil {
		return nil, err
	}
	return debts, nil
}

func (r *DebtRepository) FindActiveIncoming(lenderID string) ([]*domain.Debt, error) {
	filter := bson.M{"lender_id": lenderID, "status": bson.M{"$in": []domain.Status{domain.RequestAccepted}}}
	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	var debts []*domain.Debt
	if err = cursor.All(r.ctx, &debts); err != nil {
		return nil, err
	}
	return debts, nil
}

func (r *DebtRepository) FindActiveOutgoing(borrowerID string) ([]*domain.Debt, error) {
	filter := bson.M{"borrower_id": borrowerID, "status": bson.M{"$in": []domain.Status{domain.RequestAccepted}}}
	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	var debts []*domain.Debt
	if err = cursor.All(r.ctx, &debts); err != nil {
		return nil, err
	}
	return debts, nil
}

func (r *DebtRepository) FindHistory(userID string) ([]*domain.Debt, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"lender_id": userID},
			{"borrower_id": userID},
		},
	}
	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	var debts []*domain.Debt
	if err = cursor.All(r.ctx, &debts); err != nil {
		return nil, err
	}
	return debts, nil
}

func (r *DebtRepository) FindNetAmounts(userID string) (map[domain.Currency]float64, error) {
	// Only consider active debts
	activeStatuses := []domain.Status{
		domain.RequestAccepted,
		domain.RequestPaid,
	}

	// Debts where user is lender (they are owed money)
	lenderFilter := bson.M{
		"lender_id": userID,
		"status":    bson.M{"$in": activeStatuses},
	}
	lenderCursor, err := r.collection.Find(r.ctx, lenderFilter)
	if err != nil {
		return nil, err
	}
	var lenderDebts []*domain.Debt
	if err = lenderCursor.All(r.ctx, &lenderDebts); err != nil {
		return nil, err
	}

	// Debts where user is borrower (they owe money)
	borrowerFilter := bson.M{
		"borrower_id": userID,
		"status":      bson.M{"$in": activeStatuses},
	}
	borrowerCursor, err := r.collection.Find(r.ctx, borrowerFilter)
	if err != nil {
		return nil, err
	}
	var borrowerDebts []*domain.Debt
	if err = borrowerCursor.All(r.ctx, &borrowerDebts); err != nil {
		return nil, err
	}

	net := make(map[domain.Currency]float64)
	for _, d := range lenderDebts {
		net[d.Currency] += d.Amount
	}
	for _, d := range borrowerDebts {
		net[d.Currency] -= d.Amount
	}
	return net, nil
}

func (r *DebtRepository) FindIncomingRequests(lenderID string) ([]*domain.Debt, error) {
	filter := bson.M{"lender_id": lenderID, "status": domain.RequestPending}
	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	var debts []*domain.Debt
	if err = cursor.All(r.ctx, &debts); err != nil {
		return nil, err
	}
	return debts, nil
}
