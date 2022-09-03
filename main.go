package main

import (
	"fmt"
	"os"

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

	port := os.Getenv("PORT")
	fmt.Println(fmt.Sprintf("🚀 Server running on port %s", port))

	app.Run(fmt.Sprintf(":%s", port))
}
