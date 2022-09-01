package planetsusecase

import (
	"github.com/BaianorASR/go-star-wars/entities"
	"github.com/americanas-go/errors"
)

func (u *planetsUseCase) Create(planet *entities.Planet) (*entities.Planet, error) {
	planet, err := u.repo.Create(planet)
	if err != nil {
		return nil, errors.NewInternal(err, "error on create")
	}

	return planet, nil
}
