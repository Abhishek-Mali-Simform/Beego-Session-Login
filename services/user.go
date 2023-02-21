package services

import (
	"LoginUser/models"
	"fmt"
	"time"
)

func Create(user models.User) {
	db := GetDB()
	id, err := db.Insert(&user)
	passMsg := fmt.Sprintf("User inserted with id: %d", id)
	ErrorCheckWithSuccess("Couldn't insert new User.", passMsg, err)
}

func DummyUser() (user models.User) {
	pass := ConvertToMD5("Abhishek Mali")
	user = models.User{
		Name:      "Abhishek Mali",
		Email:     "abhishek.m@simformsolutions.com",
		Password:  pass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return
}

func Read(user *models.User, fields ...string) error {
	db := GetDB()
	err := db.Read(user, fields...)
	return err
}
