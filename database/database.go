package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_NAME = "crud_prueba"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var DB *gorm.DB

func Init() *gorm.DB {
	DB = connect()
	return DB
}

func connect() *gorm.DB {
	// var err error

	const unformatDNS = "%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local"
	dns := fmt.Sprintf(unformatDNS, DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("Error connecting to database : error= %v", err)
		return nil
	}
	return db
}
