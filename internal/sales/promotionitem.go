package sales

import (
	"context"
)

type PromoItem struct {
	SKU   string
	Price float64
	Qty   int
}

type PromotionItemService interface {
	Get(ctx context.Context, item *CartItem) (*PromoItem, error)
}
