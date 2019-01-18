package user

import (
	"github.com/jinzhu/gorm"
	"go-boilerplate/models"
)

type SQLUserRepository struct {
	DBCon *gorm.DB
}

func (sqlRepo *SQLUserRepository) GetUserByEmail(email string) (*models.User, error) {
	u := models.User{}
	err := new(error)
	return &u, *err
}

func (sqlRepo *SQLUserRepository) GetUserById(id int) (*models.User, error) {
	u := models.User{Id:id}
	err := new(error)
	return &u, *err
}

func GetNewSQLUserRepository(db *gorm.DB) UserRepository {
	return &SQLUserRepository{DBCon:db}
}
