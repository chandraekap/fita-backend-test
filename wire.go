// +build wireinject

package main

import (
	"github.com/chandraekap/fita-backend-test/graph"
	"github.com/chandraekap/fita-backend-test/internal/inventory"
	inventorymock "github.com/chandraekap/fita-backend-test/internal/inventory/mock"
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
)

var inventorySet = wire.NewSet(
	inventorymock.InitData,
	wire.FieldsOf(new(*inventorymock.Data), "Items", "Stocks"),

	inventorymock.NewItemRepository,
	wire.Bind(new(inventory.ItemRepository), new(*inventorymock.ItemRepository)),

	inventorymock.NewStockRepository,
	wire.Bind(new(inventory.StockRepository), new(*inventorymock.StockRepository)),
)

func BuildResolver() *graph.Resolver {
	wire.Build(
		salesSet,
		inventorySet,
		graph.NewResolver,
	)
	return &graph.Resolver{}
}
