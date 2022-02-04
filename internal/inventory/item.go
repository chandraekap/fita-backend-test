package inventory

import "context"

type Item struct {
	SKU          string
	Name         string
	Price        float64
	InventoryQty int
}

type ItemRepository interface {
	FindAll(ctx context.Context) ([]*Item, error)
	FindBySKU(ctx context.Context, sku string) (*Item, error)
	UpdateStock(ctx context.Context, sku string, stock int) error
}
