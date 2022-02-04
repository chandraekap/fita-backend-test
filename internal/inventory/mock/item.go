package mock

import (
	"context"
	"errors"

	"github.com/chandraekap/fita-backend-test/internal/inventory"
)

type ItemRepository struct {
	data []*inventory.Item
}

func NewItemRepository(
	data []*inventory.Item,
) *ItemRepository {
	return &ItemRepository{
		data: data,
	}
}

func (repo *ItemRepository) FindBySKU(ctx context.Context, sku string) (*inventory.Item, error) {
	for _, itemData := range repo.data {
		if itemData.SKU == sku {
			return itemData, nil
		}
	}

	return nil, errors.New("Item not found.")
}

func (repo *ItemRepository) UpdateStock(ctx context.Context, sku string, stock int) error {
	for i, itemData := range repo.data {
		if itemData.SKU == sku {
			repo.data[i].InventoryQty = stock
			return nil
		}
	}

	return errors.New("Failed to insert Item.")
}
