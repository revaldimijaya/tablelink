package usecase

import (
	"errors"
	"github/revaldimijaya/tablelink/model"
	"github/revaldimijaya/tablelink/repository"
)

type IngredientUsecase struct {
	repo *repository.IngredientRepository
}

func NewIngredientUsecase(repo *repository.IngredientRepository) *IngredientUsecase {
	return &IngredientUsecase{repo: repo}
}

func (u *IngredientUsecase) GetAll(pagination int, offset int) ([]model.Ingredient, error) {
	return u.repo.GetAll(pagination, offset)
}

func (u *IngredientUsecase) Create(ingredient model.Ingredient) error {
	existing, _ := u.repo.GetAll(1, 0)
	for _, ing := range existing {
		if ing.Name == ingredient.Name {
			return errors.New("ingredient name must be unique")
		}
	}
	return u.repo.Create(ingredient)
}

func (u *IngredientUsecase) Update(ingredient model.Ingredient) error {
	return u.repo.Update(ingredient)
}

func (u *IngredientUsecase) Delete(uuid string) error {
	return u.repo.Delete(uuid)
}
