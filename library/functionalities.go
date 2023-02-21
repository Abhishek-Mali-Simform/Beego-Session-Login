package library

import (
	"LoginUser/models"
	"LoginUser/services"
	"errors"
)

func AuthenticateUser(username, password string) (user *models.User, err error) {
	msg := "Invalid Email or Password"
	hash := services.ConvertToMD5(password)
	user = &models.User{Email: username}
	err = services.Read(user, "Email")
	if err != nil {
		if err.Error() == "<QuerySeter> no row found" {
			err = errors.New(msg)
		}
		return user, err
	} else if user.Id < 1 {
		return user, errors.New(msg)
	} else if user.Password != hash {
		return user, errors.New(msg)
	} else {
		return user, nil
	}
}
