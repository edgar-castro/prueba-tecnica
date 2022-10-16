package models

import (
	"fmt"
	"go-test-api/database"
)

type User struct {
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"uniqueIndex;not null"`
}

type UserFull struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type SingUpInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SingInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user *User) (err error) {
	err = database.DB.Create(&user).Error
	return err
}

func ExistEmail(email string) bool {
	var u User
	result := database.DB.First(&u, "email = ?", email)
	if result.RowsAffected > 0 {
		fmt.Println("Si esta registrsdo jeje")
		return true

	}
	return false
}

func GetByEmail(email string) string {
	var u User
	err := database.DB.Select("password").First(&u, "email = ?", email)
	if err != nil {
		return u.Password
	}
	return ""
}

func GetUserByEmail(email string) UserFull {
	var result UserFull
	database.DB.Raw("SELECT id, name, email, password FROM user WHERE email = ?", email).Scan(&result)
	return result

}
