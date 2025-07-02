package repository

import (
	"brolend/domain"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	dbc *mongo.Collection
	ctx context.Context
}

func NewUserRepository(dbc *mongo.Collection, ctx context.Context) domain.UserRepository {
	return &userRepository{
		dbc: dbc,
		ctx: ctx,
	}
}

func (ur *userRepository) Create(user *domain.User) (string, error) {
	iu, err := ur.dbc.InsertOne(ur.ctx, user)
	if err != nil {
		return "", err
	}
	id, ok := iu.InsertedID.(interface{ String() string })
	if ok {
		return id.String(), nil
	}
	return "", errors.New("failed to convert inserted ID to string")
}

func (ur *userRepository) FindByID(id primitive.ObjectID) (*domain.User, error) {
	var user domain.User
	err := ur.dbc.FindOne(ur.ctx, bson.D{{Key: "_id", Value: id}}).Decode(&user)
	if err != nil {
		return &domain.User{}, errors.New("User not found")
	}
	return &user, nil
}

func (ur *userRepository) FindByUserID(id string) (*domain.User, error) {
	var user domain.User
	err := ur.dbc.FindOne(ur.ctx, bson.D{{Key: "user_id", Value: id}}).Decode(&user)
	if err != nil {
		return &domain.User{}, errors.New("User not found")
	}
	return &user, nil
}

func (ur *userRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	fmt.Println("username", username)
	err := ur.dbc.FindOne(ur.ctx, bson.D{{Key: "username", Value: username}}).Decode(&user)
	if err != nil {
		return &domain.User{}, errors.New("User not found")
	}
	return &user, nil
}

func (ur *userRepository) Update(user domain.User) error {
	filter := bson.D{bson.E{Key: "_id", Value: user.ID}}
	updatedFields := bson.D{}
	if user.Username != "" {
		updatedFields = append(updatedFields, bson.E{Key: "username", Value: user.Username})
	}
	if user.Password != "" {
		updatedFields = append(updatedFields, bson.E{Key: "password", Value: user.Password})
	}
	if user.FirstName != "" {
		updatedFields = append(updatedFields, bson.E{Key: "first_name", Value: user.FirstName})
	}
	if user.LastName != "" {
		updatedFields = append(updatedFields, bson.E{Key: "last_name", Value: user.LastName})
	}

	update := bson.D{
		bson.E{
			Key: "$set", Value: updatedFields},
	}

	_, err := ur.dbc.UpdateOne(ur.ctx, filter, update)

	if err != nil {
		return errors.New("failed to update user!")
	}
	return nil
}

func (ur *userRepository) DeleteUser(objID primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: objID}}
	err := ur.dbc.FindOneAndDelete(ur.ctx, filter)
	if err.Err() != nil {
		return errors.New("Failed to delete user")
	}
	return nil
}
