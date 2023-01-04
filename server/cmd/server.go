package main

import (
	"loquacious/server/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//Initialise gin server
	r := gin.Default()
	routes.InitRoutes(r)

	r.Run(":8080")
}
