package usecase

import (
	"errors"
	"github/revaldimijaya/tablelink/model"
	"github/revaldimijaya/tablelink/repository"
)

type ItemUsecase struct {
	itemRepo    *repository.ItemRepository
	itemIngRepo *repository.ItemIngredientRepository
}

func NewItemUsecase(itemRepo *repository.ItemRepository, itemIngRepo *repository.ItemIngredientRepository) *ItemUsecase {
	return &ItemUsecase{itemRepo: itemRepo, itemIngRepo: itemIngRepo}
}

func (u *ItemUsecase) GetAll(pagination int, offset int) ([]model.Item, error) {
	return u.itemRepo.GetAll(model.Filter{
		Pagination: pagination,
		Offset:     offset,
	})
}

func (u *ItemUsecase) Create(item model.CreateItemRequest) error {
	existing, err := u.itemRepo.GetAll(model.Filter{
		Name: item.Name,
	})
	if err != nil {
		return err
	}
	if len(existing) != 0 {
		return errors.New("Name must be unique")
	}
	uuidItem, err := u.itemRepo.Create(item.Item)
	if err != nil {
		return err
	}

	for _, uuid := range item.IngredientsUUID {
		itemIng := model.ItemIngredient{ItemUUID: uuidItem, IngredientUUID: uuid}
		if err := u.itemIngRepo.Create(itemIng); err != nil {
			return err
		}
	}
	return nil
}

func (u *ItemUsecase) Update(item model.Item) error {
	existing, err := u.itemRepo.GetAll(model.Filter{
		Name: item.Name,
	})
	if err != nil {
		return err
	}
	if len(existing) != 0 {
		return errors.New("Name must be unique")
	}
	return u.itemRepo.Update(item)
}

func (u *ItemUsecase) Delete(uuid string) error {
	return u.itemRepo.Delete(uuid)
}
