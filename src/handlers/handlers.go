package handlers

import (
	create_user_handler "golang-gin-rest-api/src/handlers/create-user-handler"
	get_users_handler "golang-gin-rest-api/src/handlers/get-users-handler"
	"golang-gin-rest-api/src/repositories"
)

type Container struct {
	CreateUser create_user_handler.CreateUserHandler
	GetUsers   get_users_handler.GetUsersHandler
	// User       user_repository.UserRepository
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	// Log        logger.Logger
}

// New cria uma nova instancia dos repositórios
func New(opts Options) *Container {
	repos := repositories.New(repositories.Options{})
	return &Container{CreateUser: create_user_handler.New(repos), GetUsers: get_users_handler.New(repos)}
}
