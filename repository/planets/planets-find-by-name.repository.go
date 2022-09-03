package planetsrepository

import (
	"context"
	"fmt"

	"github.com/BaianorASR/go-star-wars/entities"
	"github.com/americanas-go/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *planetsRepository) FindByName(query string) (*entities.Planet, error) {
	var planet entities.Planet

	a := r.client.FindOne(context.TODO(), bson.D{{Key: "name", Value: query}})

	if err := a.Decode(&planet); err != nil {
		fmt.Println(err)
		return nil, errors.NotFoundf("planet not found")
	}

	return &planet, nil
}
