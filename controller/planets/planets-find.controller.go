package planetscontroller

import (
	"context"
	"net/http"

	"github.com/BaianorASR/go-star-wars/database"
	planetsrepository "github.com/BaianorASR/go-star-wars/repository/planets"
	planetsusecase "github.com/BaianorASR/go-star-wars/usecase/planets"
	"github.com/americanas-go/errors"
	"github.com/gin-gonic/gin"
)

// Find controller calls the usecase to find planets
func Find(c *gin.Context) {
	db := database.ConnectDB()
	defer db.Disconnect(context.TODO())

	repo := planetsrepository.New(db)
	useCase := planetsusecase.New(repo)

	planets, err := useCase.Find()
	if err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "No planets found",
				"status":  http.StatusNotFound,
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":  err.Error(),
				"status": http.StatusInternalServerError,
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"planets": planets,
	})
}
