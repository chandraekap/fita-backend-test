package sales

import (
	"context"
)

type CheckoutSummary struct {
	Items          []*CartItem
	DiscountAmount float64
	TotalAmount    float64
}

type CheckoutService interface {
	Checkout(ctx context.Context, clientID int) (*CheckoutSummary, error)
}

type checkoutService struct {
	cartRepository       CartRepository
	promotionRepository  PromotionRepository
	promotionItemFactory PromotionItemFactory
}

func NewCheckoutService(
	cartRepository CartRepository,
	promotionRepository PromotionRepository,
	promotionItemFactory PromotionItemFactory,
) CheckoutService {
	return &checkoutService{
		cartRepository:       cartRepository,
		promotionRepository:  promotionRepository,
		promotionItemFactory: promotionItemFactory,
	}
}

func (service *checkoutService) Checkout(ctx context.Context, clientID int) (*CheckoutSummary, error) {
	checkoutItems := []*CartItem{}
	availablePromoBySKU := map[string]*PromoItem{}
	totalAmount := 0.0

	cart, err := service.cartRepository.FindByClientID(ctx, clientID)
	if err != nil {
		return nil, err
	}

	for _, cartItem := range cart.Items {
		promo, _ := service.promotionRepository.FindBySKU(ctx, cartItem.SKU)

		if promo != nil {
			promotionItemService := service.promotionItemFactory.GetService(ctx, promo.RuleType)
			promoItem, err := promotionItemService.Get(ctx, cartItem)
			if err != nil && err != ErrPromotionItemNotFound {
				return nil, err
			}

			if cartItem.SKU != promoItem.SKU {
				availablePromoBySKU[promoItem.SKU] = promoItem
			} else if cartItem.Price != promoItem.Price {

				cartItem = &CartItem{
					SKU:   cartItem.SKU,
					Name:  cartItem.Name,
					Qty:   cartItem.Qty - promoItem.Qty,
					Price: cartItem.Price,
				}

				if cartItem.Qty > 0 {
					checkoutItems = append(checkoutItems, cartItem)
				}

				checkoutItems = append(checkoutItems, &CartItem{
					SKU:   cartItem.SKU,
					Name:  cartItem.Name,
					Qty:   promoItem.Qty,
					Price: promoItem.Price,
				})

				continue
			}
		}

		checkoutItems = append(checkoutItems, &CartItem{
			SKU:   cartItem.SKU,
			Name:  cartItem.Name,
			Qty:   cartItem.Qty,
			Price: cartItem.Price,
		})
	}

	for _, checkoutItem := range checkoutItems {
		if promoItem, ok := availablePromoBySKU[checkoutItem.SKU]; ok {

			diffPromoQty := checkoutItem.Qty - promoItem.Qty

			if diffPromoQty > 0 {
				checkoutItem.Qty = diffPromoQty
				checkoutItems = append(checkoutItems, &CartItem{
					SKU:   checkoutItem.SKU,
					Name:  checkoutItem.Name,
					Qty:   promoItem.Qty,
					Price: promoItem.Price,
				})
			} else {
				checkoutItem.Price = promoItem.Price
			}

		}
		totalAmount += (checkoutItem.Price * float64(checkoutItem.Qty))
	}

	err = service.cartRepository.Delete(ctx, cart)
	if err != nil {
		return nil, err
	}

	return &CheckoutSummary{
		Items:       checkoutItems,
		TotalAmount: totalAmount,
	}, nil
}
