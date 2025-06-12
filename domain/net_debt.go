package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NetDebt struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	User1ID  string             `bson:"userId" json:"userId"`
	User2ID  string             `bson:"user2Id" json:"user2Id"`
	Amount   float64            `bson:"amount" json:"amount"`
	Currency Currency           `bson:"currency" json:"currency"`
	Date     time.Time          `bson:"date" json:"date"`
}
