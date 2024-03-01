package controller

import (
	"log"
	"net/http"

	"api-gin/src/models/user"
	repoUser "api-gin/src/repository/user"
	serviceUser "api-gin/src/services/user"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	users, err := repoUser.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Get Users Error", "user": nil})
		return
	}

	log.Println("users :", users)

	c.JSON(http.StatusOK, gin.H{"message": "Get All Users", "user": users})
	return
}

func GetByUserID(c *gin.Context) {
	UserID := c.Param("id")

	user, err := repoUser.GetByID(UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Get User Error", "user": nil})
		return
	}

	log.Println("users :", user)
	c.JSON(http.StatusOK, gin.H{"message": "Get User Data", "user": user})
	return
}

func CreateUser(c *gin.Context) {
	var importUser user.CreateUserInput

	// Bind request body to CreateUserRequest struct and perform validation
	if err := c.ShouldBindJSON(&importUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check User Name
	user, _ := repoUser.GetByUserName(importUser.UserName)
	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "UserName Is Duplicate", "user": nil})
		return
	}

	// Convert Body
	log.Println("importUser:", importUser)
	bodyUser := serviceUser.MapBodyCreateUser(importUser)

	// Save User In DB
	err := repoUser.CreateUser(bodyUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Create Users Error", "user": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Create One User", "user": bodyUser})
	return
}

func UpdateUser(c *gin.Context) {
	var updateUser user.UpdateUserInput

	// Bind validation
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check ID
	log.Println("ID :", updateUser.ID)
	user, _ := repoUser.GetByID(updateUser.ID)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID Not Found", "user": nil})
		return
	}

	// Convert Body
	log.Println("updateUser:", updateUser)
	bodyUser := serviceUser.MapBodyUpdateUser(updateUser, user)

	// Save User In DB
	err := repoUser.UpdateByID(updateUser.ID, bodyUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Update Users Error", "user": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update User Success", "user": bodyUser})
	return
}

func DeleteByUserID(c *gin.Context) {
	UserID := c.Param("id")

	// Delete User In DB
	err := repoUser.DeleteByID(UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Delete User Error", "user": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete User Success", "id": UserID})
	return
}
