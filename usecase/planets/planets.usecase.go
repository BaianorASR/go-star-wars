package planetsusecase

import (
	planetsrepository "github.com/BaianorASR/go-star-wars/repository/planets"
)

type planetsUseCase struct {
	repo planetsrepository.IPlanetsRepository
}

func New(r planetsrepository.IPlanetsRepository) planetsUseCase {
	return planetsUseCase{
		repo: r,
	}
}
