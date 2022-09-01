package planetsrepository

import (
	"github.com/BaianorASR/go-star-wars/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type planetsRepository struct {
	client *mongo.Collection
}

type IPlanetsRepository interface {
	Find(query string) (*[]entities.Planet, error)
	FindById(id primitive.ObjectID) (*entities.Planet, error)
	DeleteOne(id primitive.ObjectID) error
	Create(planet *entities.Planet) (*entities.Planet, error)
}

func New(db *mongo.Client) *planetsRepository {
	return &planetsRepository{
		client: db.Database("go-americanas-teste").Collection("planets"),
	}
}
