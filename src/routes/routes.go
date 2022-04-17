package routes

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func New() routes {
	return registerRoutes()
}

func registerRoutes() routes {
	r := routes{
		router: gin.Default(),
	}
	api := r.router.Group("/api")

	// call all routes
	r.addUsersRoutes(api)

	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run()
}
