package usecase

import (
	"github/revaldimijaya/tablelink/repository"
)

type ItemIngredientUsecase struct {
	repo *repository.ItemIngredientRepository
}

func NewItemIngredientUsecase(repo *repository.ItemIngredientRepository) *ItemIngredientUsecase {
	return &ItemIngredientUsecase{repo: repo}
}

func (u *ItemIngredientUsecase) Delete(itemUUID string, ingredientUUID string) error {
	return u.repo.Delete(itemUUID, ingredientUUID)
}
