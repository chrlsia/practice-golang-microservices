package services

import "github.com/chrlsia/practice-golang-microservices/mvc/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
