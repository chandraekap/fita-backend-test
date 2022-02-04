package inventory

import "context"

type Stock struct {
	ID     int
	SKU    string
	Reason string
	Qty    int
}

type StockRepository interface {
	FindLastBySKU(ctx context.Context, sku string) (*Stock, error)
	Insert(ctx context.Context, stock *Stock) error
}
