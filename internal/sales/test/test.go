package test

import "github.com/chandraekap/fita-backend-test/internal/sales"

type TestContainer struct {
	CheckoutService                sales.CheckoutService
	PromotionItemBuy2Free1         *sales.PromotionItemBuy2Free1
	PromotionItemDiscount10Percent *sales.PromotionItemDiscount10Percent
}

func NewTestContainer(
	checkoutService sales.CheckoutService,
	promotionItemBuy2Free1 *sales.PromotionItemBuy2Free1,
	promotionItemDiscount10Percent *sales.PromotionItemDiscount10Percent,
) *TestContainer {
	return &TestContainer{
		CheckoutService:                checkoutService,
		PromotionItemBuy2Free1:         promotionItemBuy2Free1,
		PromotionItemDiscount10Percent: promotionItemDiscount10Percent,
	}
}
