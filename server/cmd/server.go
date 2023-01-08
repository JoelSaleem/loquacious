package main

import (
	"loquacious/server/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

// todo: change package to include github

type Specification struct {
	ApiKey    string `default:"foobar" split_words:"true"`    // todo: make required
	RedisHost string `default:"localhost" split_words:"true"` // todo: make required
}

func main() {
	var s Specification
	err := envconfig.Process("server", &s)
	if err != nil {
		panic(err.Error())
	}

	//Initialise gin server
	r := gin.Default()
	routes.InitRoutes(r, s.RedisHost)

	r.Run(":8080")
}
