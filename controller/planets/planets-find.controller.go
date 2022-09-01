package planetscontroller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/BaianorASR/go-star-wars/database"
	planetsrepository "github.com/BaianorASR/go-star-wars/repository/planets"
	planetsusecase "github.com/BaianorASR/go-star-wars/usecase/planets"
	"github.com/gin-gonic/gin"
)

// Find controller calls the usecase to find planets
func Find(c *gin.Context) {
	db := database.ConnectDB()

	names, err := db.ListDatabaseNames(context.TODO(), nil)
	fmt.Println(names)

	repo := planetsrepository.New(db)
	useCase := planetsusecase.New(repo)

	query := c.Query("query")

	planets, err := useCase.Find(query)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"planets": planets,
	})
}
