package controller

import (
	mdw "api-gin/src/middleware"
	"api-gin/src/models/auth"
	repoUser "api-gin/src/repository/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var inputLogIn auth.UserLogin

	// Bind Validation
	if err := c.ShouldBindJSON(&inputLogIn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check User Name
	user, _ := repoUser.GetByUserName(inputLogIn.UserName)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User Not Found", "data": nil})
		return
	}

	// Check Password
	checkPassword := mdw.ComparePassword(inputLogIn.Password, user.Password)
	if !checkPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password Is Valid", "data": nil})
		return
	}

	// Gen Token
	token, err := mdw.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Genarate Token Error", "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Login Success", "token": token})
	return
}
