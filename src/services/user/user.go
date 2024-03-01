package user

import (
	"api-gin/src/models/user"
	"strings"
	"time"
)

func MapBodyCreateUser(dataInput user.CreateUserInput) *user.Users {
	return &user.Users{
		UserName:  dataInput.UserName,
		Gender:    strings.ToUpper(dataInput.Gender),
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func MapBodyUpdateUser(dataInput user.UpdateUserInput, userData *user.Users) *user.Users {
	return &user.Users{
		UserName:  userData.UserName,
		Gender:    strings.ToUpper(dataInput.Gender),
		Active:    dataInput.Active,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: time.Now(),
	}
}
