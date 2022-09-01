package planetsusecase

import (
	planetsrepository "github.com/BaianorASR/go-star-wars/repository/planets"
)

type planetsUseCase struct {
	repo planetsrepository.IPlanetsRepository
}

type IPlanetsUseCase interface {
	// Find(query string) (*[]entities.Planet, error)
}

func New(r planetsrepository.IPlanetsRepository) planetsUseCase {
	return planetsUseCase{
		repo: r,
	}
}
