package user

import (
	"go-boilerplate/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id int) (*models.User, error)
}