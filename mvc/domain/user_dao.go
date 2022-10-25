package domain

import (
	"fmt"
	"github.com/chrlsia/practice-golang-microservices/mvc/utils"
	"net/http"
)

var users = map[int64]*User{
	123: {
		Id: 123, FirstName: "Federico", LastName: "Leon", Email: "myemail@gmail.com"},
}

func GetUser(UserId int64) (*User, *utils.ApplicationError) {

	if user := users[UserId]; user != nil {
		return user, nil

	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", UserId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}

}
