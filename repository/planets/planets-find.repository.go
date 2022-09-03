package planetsrepository

import (
	"context"

	"github.com/BaianorASR/go-star-wars/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Find return all planets
func (r *planetsRepository) Find() (*[]entities.Planet, error) {
	var planets []entities.Planet

	data, err := r.client.Find(context.TODO(), bson.D{}, &options.FindOptions{
		Sort: bson.D{{Key: "name", Value: 1}},
	})
	if err != nil {
		return nil, err
	}

	if err := data.All(context.TODO(), &planets); err != nil {
		return nil, err
	}

	return &planets, nil
}
