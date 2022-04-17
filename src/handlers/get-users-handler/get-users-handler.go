package get_users_handler

import (
	"golang-gin-rest-api/src/repositories"
	GetUsersUseCase "golang-gin-rest-api/src/usecases/get-users-use-case"

	"github.com/gin-gonic/gin"
)

type GetUsersHandler interface {
	Execute(c *gin.Context)
}

type handler struct {
	repositories *repositories.Container
	// log          logger.Logger
}

func New(repo *repositories.Container /*, log logger.Logger */) GetUsersHandler {
	return &handler{repositories: repo}
}

func (h *handler) Execute(c *gin.Context) {

	getUsersUseCase := GetUsersUseCase.New(h.repositories)

	// fazer factory use case get users
	payload := getUsersUseCase.Execute()

	c.JSON(200, gin.H{
		"response": payload,
	})
}
