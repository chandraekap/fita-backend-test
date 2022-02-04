package sales

import "context"

const RASPBERRY_PI_PROMO_SKU string = "234234"

type PromotionItemFreeRaspberry struct {
}

func NewPromotionItemFreeRaspberry() *PromotionItemFreeRaspberry {
	return &PromotionItemFreeRaspberry{}
}

func (service *PromotionItemFreeRaspberry) Get(ctx context.Context, item *CartItem) (*PromoItem, error) {
	return &PromoItem{
		SKU:   RASPBERRY_PI_PROMO_SKU,
		Price: 0,
		Qty:   1,
	}, nil
}
