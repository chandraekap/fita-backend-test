package sales

import (
	"context"
	"errors"
)

type PromotionItemBuy2Free1 struct {
}

func NewPromotionItemBuy2Free1() *PromotionItemBuy2Free1 {
	return &PromotionItemBuy2Free1{}
}

func (service *PromotionItemBuy2Free1) Get(ctx context.Context, item *CartItem) (*PromoItem, error) {
	if item.Qty < 3 {
		return nil, errors.New("Quantity must be grater than 3.")
	}

	count := item.Qty / 3

	return &PromoItem{
		SKU:   item.SKU,
		Price: 0,
		Qty:   count,
	}, nil
}
