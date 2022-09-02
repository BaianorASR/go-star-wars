package planetscontroller

import (
	"context"
	"net/http"

	"github.com/BaianorASR/go-star-wars/database"
	planetsrepository "github.com/BaianorASR/go-star-wars/repository/planets"
	planetsusecase "github.com/BaianorASR/go-star-wars/usecase/planets"
	"github.com/americanas-go/errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindByID controller calls the usecase to find a planet by id
func FindById(c *gin.Context) {
	var err error

	db := database.ConnectDB()
	defer db.Disconnect(context.TODO())

	repo := planetsrepository.New(db)
	usecase := planetsusecase.New(repo)

	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid hex id",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var objId primitive.ObjectID
	if objId, err = primitive.ObjectIDFromHex(id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	planet, err := usecase.FindById(objId)
	if err != nil {
		if errors.IsNotFound(err) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"planet": planet})

}
