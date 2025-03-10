package user_service

import (
	"ecomm-crud/pkg/models/user-models"
)


func GetUsers() []models.User {

	var users []models.User
	var user models.User
	user.ID = 18
	user.Name = "test"
	users = append(users, user)

	return users
}