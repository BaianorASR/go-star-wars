package main

import (
	"github.com/BaianorASR/go-star-wars/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	gin.Recovery()
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	router := app.Group("/api/v1")
	routes.AddRoutes(router)
	println("1")

	if err := app.Run(":3001"); err != nil {
		panic(err)
	}
}
