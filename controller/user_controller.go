package controller

import (
	"brolend/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	uu domain.UserUsecase
}

func NewUserController(uu domain.UserUsecase) *UserController {
	return &UserController{
		uu: uu,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return 
	}

	userID, err, token := uc.uu.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "userID": userID, "token": token})
	return 

}

func (uc *UserController) Update(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := uc.uu.Update(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
	return 
}

func (uc *UserController) Login(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")

	userID, err, token := uc.uu.Login(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "userID": userID, "token": token})
	return 
}

func (uc *UserController) FindUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := uc.uu.Search(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}
	c.JSON(http.StatusOK, user)
	return 
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	userIDParam := c.Param("username")
	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid user ID")
		return 
	}

	err = uc.uu.Delete(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})

	return 
}
