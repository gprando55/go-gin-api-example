package create_user_use_case

import (
	"golang-gin-rest-api/src/entities"
	"golang-gin-rest-api/src/repositories"
	"time"

	"github.com/google/uuid"
)

type CreateUserUseCase interface {
	Execute(name string, age int16) entities.User
}

type usecase struct {
	repositories *repositories.Container
	// log          logger.Logger
}

func New(repo *repositories.Container /*, log logger.Logger */) CreateUserUseCase {
	return &usecase{repositories: repo}
}

func (uc *usecase) Execute(name string, age int16) entities.User {

	id := uuid.New().String()

	user := entities.User{ID: id, Name: name, Age: age, CreatedAt: time.Now()}

	err := uc.repositories.User.Create(user)

	if err != nil {
		panic("Error on create new user")
	}

	return user
}
