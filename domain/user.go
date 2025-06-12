package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        string `bson:"_id" json:"id"`
	UserID    string `bson:"user_id" json:"user_id"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Username  string `bson:"username" json:"username"`
	Password  string `bson:"password" json:"password"`
}

type UserRepository interface {
	Create(user *User) (string, error)
	FindByID(id primitive.ObjectID) (*User, error)
	FindByUserID(id string) (*User, error)
	FindByUsername(username string) (*User, error)
	Update(user *User) error
	DeleteUser(objID primitive.ObjectID) error
}
