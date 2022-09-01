package planetsusecase

import "go.mongodb.org/mongo-driver/bson/primitive"

func (u *planetsUseCase) DeleteOne(objId primitive.ObjectID) error {
	return u.repo.DeleteOne(objId)
}
