package database

import (
	"fmt"
	"log"
	"test-gonic/config/db_config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	var ConnectionError error

	if db_config.DB_DRIVER == "mysql" {
		dsnMysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_HOST, db_config.DB_PORT, db_config.DB_NAME)

		DB, ConnectionError = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{})
	}

	if db_config.DB_DRIVER == "pgsql" {
		panic("Cant connect to database PQSQL")
	}

	if ConnectionError != nil {
		panic("Cant connect to database")
	}

	log.Println("Connect database")

}
