package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-boilerplate/config"
)

/*
We keep a global reference to the sql DB in our Data store class. 
This is a thread safe DB pool and this will be used across packages
*/
var DBCon *gorm.DB

/*
Initialize the DB string and connects to the database. 
The global DBCon will be assigned in this
*/
func Connect() bool {
	dbConnectString := config.GetConfig("DB_USER_NAME") + ":" + config.GetConfig("DB_PASSWORD") + "@" +
		config.GetConfig("DB_URL") + "/" + config.GetConfig("DB_NAME")

	var err error
	DBCon, err = gorm.Open("mysql", dbConnectString + "?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to db")
	}

	return true
}

func GetDBConnection() *gorm.DB {
	return DBCon
}
