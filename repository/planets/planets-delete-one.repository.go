package planetsrepository

import (
	"context"

	"github.com/americanas-go/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteOne deletes a planet by id
func (r *planetsRepository) DeleteOne(objId primitive.ObjectID) error {

	result, err := r.client.DeleteOne(context.TODO(), bson.M{"_id": objId})
	if err != nil {
		return errors.NewInternal(err, "error on delete one")
	}

	if result.DeletedCount == 0 {
		return errors.NewNotFound(err, "planet not found")
	}

	return nil
}
