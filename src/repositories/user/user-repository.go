package user_repository

import (
	"golang-gin-rest-api/src/entities"
)

type UserRepository interface {
	Create(newUser entities.User) error
	GetAll() ([]entities.User, error)
}
