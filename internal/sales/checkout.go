package sales

import (
	"context"

	apperrors "github.com/chandraekap/fita-backend-test/internal/errors"
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
			_, ok := err.(*apperrors.NotFoundError)
			if err != nil && !ok {
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
			checkoutItem.Price = promoItem.Price
		}
		totalAmount += (checkoutItem.Price * float64(checkoutItem.Qty))
	}

	return &CheckoutSummary{
		Items:       checkoutItems,
		TotalAmount: totalAmount,
	}, nil
}
