package main

import (
	"fmt"

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

	fmt.Println("🚀 Server running on port 3001")
	app.Run(":3001")
}
