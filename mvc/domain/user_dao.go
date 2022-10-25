package domain

import (
	"errors"
	"fmt"
)

var users = map[int64]*User{
	123: {
		Id: 1, FirstName: "Federico", LastName: "Leon", Email: "myemail@gmail.com"},
}

func GetUser(UserId int64) (*User, error) {

	if user := users[UserId]; user != nil {
		return user, nil

	}
	return nil, errors.New(fmt.Sprintf("user %v was not found", UserId))

}
