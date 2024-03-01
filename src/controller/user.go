package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"api-gin/src/models/user"
	repoUser "api-gin/src/repository/user"
	serviceUser "api-gin/src/services/user"

	"api-gin/src/db/redis"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {

	cacheVal := redis.GetValue("allUsers")
	if cacheVal != "" {
		// Assuming jsonString is the JSON string you want to convert
		var dataUser []user.Users
		err := json.Unmarshal([]byte(cacheVal), &dataUser)
		if err != nil {
			// Handle error
			fmt.Println("Error unmarshaling JSON:", err)
			return
		}

		log.Println("*****  Cache Value *****")

		c.JSON(http.StatusOK, gin.H{"message": "Get All Users", "user": dataUser})
		return
	}

	users, err := repoUser.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Get Users Error", "user": nil})
		return
	}

	// Set Value in Redis
	usersJSON, err := json.Marshal(users)
	if err != nil {
		// Handle error
		fmt.Println("Error converting users to JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Get Users Error", "user": nil})
		return
	}
	usersString := string(usersJSON)
	redis.SetKey("allUsers", usersString)
	log.Println("*****  Set Value *****")

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

	// Clear Cache
	redis.ClearCache("allUsers")

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

	c.JSON(http.StatusCreated, gin.H{"message": "Create User Success", "user": nil})
	return
}

func UpdateUser(c *gin.Context) {
	var updateUser user.UpdateUserInput

	// Clear Cache
	redis.ClearCache("allUsers")

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

	// Clear Cache
	redis.ClearCache("allUsers")

	// Delete User In DB
	err := repoUser.DeleteByID(UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Delete User Error", "user": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete User Success", "id": UserID})
	return
}
