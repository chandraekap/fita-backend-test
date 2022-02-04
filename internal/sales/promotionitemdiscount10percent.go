package sales

import (
	"context"
)

type PromotionItemDiscount10Percent struct {
}

func NewPromotionItemDiscount10Percent() *PromotionItemDiscount10Percent {
	return &PromotionItemDiscount10Percent{}
}

func (service *PromotionItemDiscount10Percent) Get(ctx context.Context, item *CartItem) (*PromoItem, error) {
	if item.Qty < 3 {
		return nil, ErrPromotionItemNotFound
	}

	discount := item.Price * 0.1

	return &PromoItem{
		SKU:   item.SKU,
		Price: item.Price - discount,
		Qty:   item.Qty,
	}, nil
}
