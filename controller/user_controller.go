package controller

import (
	"brolend/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userController struct {
	uu domain.UserUsecase
}

func NewUserController(uu domain.UserUsecase) *userController {
	return &userController{
		uu: uu,
	}
}

func (uc *userController) Register(c *gin.Context) error {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	userID, err, token := uc.uu.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "userID": userID, "token": token})
	return nil

}

func (uc *userController) Update(c *gin.Context) error {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	err := uc.uu.Update(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
	return nil
}

func (uc *userController) Login(c *gin.Context) error {
	username := c.Param("username")
	password := c.Param("password")

	userID, err, token := uc.uu.Login(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "userID": userID, "token": token})
	return nil
}

func (uc *userController) FindUserByUsername(c *gin.Context) error {
	username := c.Param("username")
	user, err := uc.uu.Search(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	c.JSON(http.StatusOK, user)
	return nil
}

func (uc *userController) DeleteUser(c *gin.Context) error {
	userIDParam := c.Param("username")
	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid user ID")
		return err
	}

	err = uc.uu.Delete(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})

	return nil
}
