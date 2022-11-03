package repository

import "golang-jwt-example/pkg/domain/entity"

type IUserRepository interface {
	GetUser(userID int) (*entity.User, error)
}