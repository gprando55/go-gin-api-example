package get_users_use_case

import (
	"golang-gin-rest-api/src/entities"
	"golang-gin-rest-api/src/repositories"
)

type GetUsersUseCase interface {
	Execute() []entities.User
}

type usecase struct {
	repositories *repositories.Container
	// log          logger.Logger
}

func New(repo *repositories.Container /*, log logger.Logger */) GetUsersUseCase {
	return &usecase{repositories: repo}
}

func (uc *usecase) Execute() []entities.User {

	users, err := uc.repositories.User.GetAll()

	if err != nil {
		panic("Error on get all users")
	}

	return users
}
