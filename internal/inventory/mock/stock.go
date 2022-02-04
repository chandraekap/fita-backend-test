package mock

import (
	"context"
	"errors"

	"github.com/chandraekap/fita-backend-test/internal/inventory"
)

type StockRepository struct {
	data []*inventory.Stock
}

func NewStockRepository(
	data []*inventory.Stock,
) *StockRepository {
	return &StockRepository{
		data: data,
	}
}

func (repo *StockRepository) FindLastBySKU(ctx context.Context, sku string) (*inventory.Stock, error) {
	for _, stockData := range repo.data {
		if stockData.SKU == sku {
			return stockData, nil
		}
	}

	return nil, errors.New("Item not found.")
}

func (repo *StockRepository) Insert(ctx context.Context, stock *inventory.Stock) error {
	repo.data = append(repo.data, stock)

	return errors.New("Failed to insert Stock.")
}
