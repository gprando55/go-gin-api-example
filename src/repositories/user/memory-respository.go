package user_repository

import (
	"golang-gin-rest-api/src/entities"
)

var usersMemory []entities.User

type repoMemory struct {
	users []entities.User
}

func NewMemoryRepository() UserRepository {

	return &repoMemory{users: usersMemory}
}

func (repo *repoMemory) Create(newUser entities.User) error {

	repo.users = append(repo.users, newUser)

	return nil
}

func (repo *repoMemory) GetAll() ([]entities.User, error) {

	return repo.users, nil
}
