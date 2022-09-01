package planetsrepository

import (
	"context"

	"github.com/BaianorASR/go-star-wars/entities"
	"go.mongodb.org/mongo-driver/bson"
)

// Find return all planets when query is empty or planets that match with query
func (r *planetsRepository) Find(query string) (*[]entities.Planet, error) {
	println(query)

	filter := bson.D{{Key: "name", Value: bson.D{{Key: "$regex", Value: query}}}}

	// Make pagination letter
	// options := new(options.FindOptions)
	// options.SetLimit(10)
	// options.SetSkip(0)

	data, err := r.client.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var planets []entities.Planet
	if err := data.All(context.TODO(), &planets); err != nil {
		return nil, err
	}

	return &planets, nil
}
