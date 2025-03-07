package usecase

import (
	"errors"
	"fmt"
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
	return u.repo.GetAll(model.Filter{
		Pagination: pagination,
		Offset:     offset,
	})
}

func (u *IngredientUsecase) Create(ingredient model.Ingredient) error {
	fmt.Printf("%+v\n", ingredient)
	existing, _ := u.repo.GetAll(model.Filter{
		Name:       ingredient.Name,
		Pagination: 1,
		Offset:     0,
	})
	fmt.Printf("%+v\n", existing)
	if len(existing) != 0 {
		return errors.New("ingredient name must be unique")
	}
	return u.repo.Create(ingredient)
}

func (u *IngredientUsecase) Update(ingredient model.Ingredient) error {
	existing, _ := u.repo.GetAll(model.Filter{
		Name:       ingredient.Name,
		Pagination: 1,
		Offset:     0,
	})
	if len(existing) != 0 {
		return errors.New("ingredient name must be unique")
	}
	return u.repo.Update(ingredient)
}

func (u *IngredientUsecase) Delete(uuid string) error {
	return u.repo.Delete(uuid)
}
