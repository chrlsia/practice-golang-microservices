package services

import (
	"github.com/chrlsia/practice-golang-microservices/mvc/domain"
	"github.com/chrlsia/practice-golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
