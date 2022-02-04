package mock

import (
	"context"
	"errors"

	"github.com/chandraekap/fita-backend-test/internal/sales"
)

type PromotionRepository struct {
	data []*sales.Promotion
}

func NewPromotionRepository(
	data []*sales.Promotion,
) *PromotionRepository {
	return &PromotionRepository{
		data: data,
	}
}

func (repo *PromotionRepository) FindBySKU(ctx context.Context, sku string) (*sales.Promotion, error) {
	for _, promoData := range repo.data {
		if promoData.ItemSKU == sku {
			return promoData, nil
		}
	}

	return nil, errors.New("Promo not found.")
}
