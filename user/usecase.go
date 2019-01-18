package user

import (
	"go-boilerplate/models"
)

type UserUseCase struct {
	userRepo UserRepository
}

func (u *UserUseCase) IsValidUser(email string) *models.User  {
	user := &models.User{}
	return user
}

func (u *UserUseCase) GetUserById(id int) *models.User {
	user, _ := u.userRepo.GetUserById(id)
	return user
}

func GetNewUserUseCase (r UserRepository) *UserUseCase {
	return &UserUseCase{userRepo:r}
}