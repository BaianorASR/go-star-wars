package planetsrepository

import (
	"context"

	"github.com/BaianorASR/go-star-wars/entities"
	"github.com/americanas-go/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindById finds a planet by id
func (r *planetsRepository) FindById(objId primitive.ObjectID) (*entities.Planet, error) {
	var planet entities.Planet

	if err := r.client.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&planet); err != nil {
		return nil, errors.NotFoundf("planet not found")
	}

	return &planet, nil
}
