package create_user_handler

import (
	"fmt"
	"golang-gin-rest-api/src/repositories"
	CreateUserUseCase "golang-gin-rest-api/src/usecases/create-user-use-case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Name string `json:"name"`
	Age  int16  `json:"age"`
}

type CreateUserHandler interface {
	Execute(c *gin.Context)
}

type handler struct {
	repositories *repositories.Container
	// log          logger.Logger
}

func New(repo *repositories.Container /*, log logger.Logger */) CreateUserHandler {
	return &handler{repositories: repo}
}

func (h *handler) Execute(c *gin.Context) {
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error on parse request",
		})

		return
	}

	fmt.Println("request body ->", requestBody)

	createUserUseCase := CreateUserUseCase.New(h.repositories)

	// fazer factory use case get users
	payload := createUserUseCase.Execute(requestBody.Name, requestBody.Age)

	c.JSON(200, gin.H{
		"response": payload,
	})
}
