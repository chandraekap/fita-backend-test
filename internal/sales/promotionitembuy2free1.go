package sales

import (
	"context"
)

type PromotionItemBuy2Free1 struct {
}

func NewPromotionItemBuy2Free1() *PromotionItemBuy2Free1 {
	return &PromotionItemBuy2Free1{}
}

func (service *PromotionItemBuy2Free1) Get(ctx context.Context, item *CartItem) (*PromoItem, error) {
	if item.Qty < 3 {
		return nil, ErrPromotionItemNotFound
	}

	count := item.Qty / 3

	return &PromoItem{
		SKU:   item.SKU,
		Price: 0,
		Qty:   count,
	}, nil
}
