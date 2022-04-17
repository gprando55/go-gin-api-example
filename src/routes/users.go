package routes

import (
	"golang-gin-rest-api/src/handlers"

	"github.com/gin-gonic/gin"
)

func (r routes) addUsersRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	h := handlers.New(handlers.Options{})

	users.GET("/", h.GetUsers.Execute)
	users.POST("/", h.CreateUser.Execute)
}
