package controller

import (
	"brolend/domain"
	"fmt"
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
	usr := domain.User{}
	if err := c.ShouldBindJSON(&usr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	user, err, token := uc.uu.Login(usr.Username, usr.Password)
	if err != nil {
		switch err.Error() {
		case "user not found":
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		case "wrong password":
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user, "token": token})
	return
}

func (uc *UserController) FindUserByUsername(c *gin.Context) {
	username := c.Param("username")
	fmt.Println("username", username)
	user, err := uc.uu.Search(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	userIDParam := c.Param("user_id")
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
