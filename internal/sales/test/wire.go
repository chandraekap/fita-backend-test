// +build wireinject

package test

import (
	"github.com/chandraekap/fita-backend-test/internal/sales"
	salesmock "github.com/chandraekap/fita-backend-test/internal/sales/mock"
	"github.com/google/wire"
)

var salesSet = wire.NewSet(
	salesmock.InitData,
	wire.FieldsOf(new(*salesmock.Data), "Carts", "Promos"),

	salesmock.NewCartRepository,
	wire.Bind(new(sales.CartRepository), new(*salesmock.CartRepository)),

	salesmock.NewPromotionRepository,
	wire.Bind(new(sales.PromotionRepository), new(*salesmock.PromotionRepository)),

	sales.NewPromotionItemFactory,
	sales.NewCheckoutService,
	sales.NewPromotionItemBuy2Free1,
	sales.NewPromotionItemDiscount10Percent,
)

func BuildTestContainer() *TestContainer {
	wire.Build(
		salesSet,
		NewTestContainer,
	)
	return &TestContainer{}
}
