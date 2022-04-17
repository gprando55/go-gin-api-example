package main

import "golang-gin-rest-api/src/routes"

func main() {
	routes := routes.New()

	routes.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
