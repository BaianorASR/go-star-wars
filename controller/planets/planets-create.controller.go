package planetscontroller

import (
	"context"
	"net/http"

	"github.com/BaianorASR/go-star-wars/database"
	"github.com/BaianorASR/go-star-wars/entities"
	planetsrepository "github.com/BaianorASR/go-star-wars/repository/planets"
	planetsusecase "github.com/BaianorASR/go-star-wars/usecase/planets"
	"github.com/americanas-go/errors"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	db := database.ConnectDB()
	defer db.Disconnect(context.TODO())

	repo := planetsrepository.New(db)
	useCase := planetsusecase.New(repo)

	var planet *entities.Planet
	c.BindJSON(&planet)

	planet, err := useCase.Create(planet)
	if err != nil {
		if errors.IsUnauthorized(err) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":  err.Error(),
				"status": http.StatusUnauthorized,
			})
		} else if errors.IsConflict(err) {
			c.JSON(http.StatusConflict, gin.H{
				"error":  err.Error(),
				"status": http.StatusConflict,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":  err.Error(),
				"status": http.StatusInternalServerError,
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"planet": planet,
	})
}
