package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	UserID    string             `bson:"user_id" json:"user_id"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `bson:"password" json:"password"`
}

type UserRepository interface {
	Create(user *User) (string, error)
	FindByID(id primitive.ObjectID) (*User, error)
	FindByUserID(id string) (*User, error)
	FindByUsername(username string) (*User, error)
	Update(user User) error
	DeleteUser(objID primitive.ObjectID) error
}

type UserUsecase interface {
	Login(username string, password string) (*User, error, string)
	Register(user User) (string, error, string)
	Search(username string) (*User, error)
	Update(user User) error
	Delete(objId primitive.ObjectID) error
	FindByID(id primitive.ObjectID) (*User, error)
}
