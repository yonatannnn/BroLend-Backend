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
	res, err := r.collection.InsertOne(r.ctx, debt)
	if err != nil {
		return "", err
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if ok {
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
	oid, err := primitive.ObjectIDFromHex(debt.ID)
	if err != nil {
		return err
	}
	_, err = r.collection.ReplaceOne(r.ctx, bson.M{"_id": oid}, debt)
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
	filter := bson.M{"lender_id": lenderID, "status": bson.M{"$in": []domain.Status{domain.RequestPending, domain.RequestAccepted, domain.RequestPaid}}}
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
	filter := bson.M{"borrower_id": borrowerID, "status": bson.M{"$in": []domain.Status{domain.RequestPending, domain.RequestAccepted, domain.RequestPaid}}}
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
		"status": domain.Settled,
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
	// This is a simplified aggregation. For production, use MongoDB aggregation pipeline.
	debts, err := r.FindByLenderID(userID)
	if err != nil {
		return nil, err
	}
	net := make(map[domain.Currency]float64)
	for _, d := range debts {
		net[d.Currency] += d.Amount
	}
	debts, err = r.FindByBorrowerID(userID)
	if err != nil {
		return nil, err
	}
	for _, d := range debts {
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
