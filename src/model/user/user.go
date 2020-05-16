package user

import (
	"errors"
	"fmt"
)

func Login(userName, password string) (User, error) {
	if _, ok := UserMap[userName]; ok {
		if UserMap[userName].Password == password {
			fmt.Println("\nLogin Successful ")
			return UserMap[userName], nil
		}
	}

	return User{}, errors.New("Invalid Login Credentials")
}

//Signup function returns userId
func (user *User) Signup() (int64, error) {
	lastId := (int64)(len(UserMap))

	if _, ok := UserMap[user.UserName]; ok || len(user.UserName) == 0 {
		return 0, errors.New("UserName already exists or is empty ")
	}

	user.ID = lastId + 1

	UserMap[user.UserName] = *user

	return user.ID, nil
}
