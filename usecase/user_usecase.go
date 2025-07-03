package usecase

import (
	"brolend/domain"
	"brolend/infrastructure"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUsecase struct {
	userRepository  domain.UserRepository
	passwordService infrastructure.PasswordService
	jwtService      infrastructure.JWTService
}

func NewUserUsecase(ur domain.UserRepository, ps infrastructure.PasswordService, js infrastructure.JWTService) domain.UserUsecase {
	return &userUsecase{
		userRepository:  ur,
		passwordService: ps,
		jwtService:      js,
	}
}

func (u *userUsecase) Login(username string, password string) (*domain.User, error, string) {
	existingUser, err := u.userRepository.FindByUsername(username)
	if err != nil {
		return &domain.User{}, errors.New("User not found"), ""
	}

	err = u.passwordService.ComparePassword(existingUser.Password, password)
	if err != nil {
		return &domain.User{}, errors.New("Invalid password"), ""
	}

	token, err := u.jwtService.GenerateToken(*existingUser)
	if err != nil {
		return nil, errors.New("Error generating token"), ""
	}

	return existingUser, nil, token
}

func (u *userUsecase) Register(user domain.User) (string, error, string) {
	user.ID = primitive.NewObjectID()
	user.UserID = user.ID.Hex()

	p, err := u.passwordService.HashPassword(user.Password)
	if err != nil {
		return "", errors.New("Error Hashing Password"), ""
	}
	user.Password = p

	existingUser, err := u.userRepository.FindByUsername(user.Username)
	if err == nil && existingUser != nil {
		return "", errors.New("Username already exists"), ""
	}

	// Generate JWT token
	token, err := u.jwtService.GenerateToken(user)
	if err != nil {
		return "", errors.New("Error generating token"), ""
	}

	_, err = u.userRepository.Create(&user)
	if err != nil {
		return "", errors.New("Error creating user"), ""
	}

	return user.UserID, nil, token
}

func (u *userUsecase) Search(username string) (*domain.User, error) {
	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("User not found")
	}
	return user, nil
}

func (u *userUsecase) Update(updatedUser domain.User) error {
	// Find existing user by ID
	existingUser, err := u.userRepository.FindByID(updatedUser.ID)
	if err != nil || existingUser == nil {
		return errors.New("User not found")
	}

	// If password is being updated, hash it
	if updatedUser.Password != "" && updatedUser.Password != existingUser.Password {
		hashedPassword, err := u.passwordService.HashPassword(updatedUser.Password)
		if err != nil {
			return errors.New("Error hashing password")
		}
		updatedUser.Password = hashedPassword
	} else {
		updatedUser.Password = existingUser.Password
	}

	// Update user in repository
	err = u.userRepository.Update(updatedUser)
	if err != nil {
		return errors.New("Error updating user")
	}

	return nil
}

func (u *userUsecase) Delete(objId primitive.ObjectID) error {
	// Check if user exists
	existingUser, err := u.userRepository.FindByID(objId)
	if err != nil || existingUser == nil {
		return errors.New("User not found")
	}

	// Delete user
	err = u.userRepository.DeleteUser(objId)
	if err != nil {
		return errors.New("Error deleting user")
	}

	return nil
}

func (u *userUsecase) FindByID(id primitive.ObjectID) (*domain.User, error) {
	user, err := u.userRepository.FindByID(id)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
