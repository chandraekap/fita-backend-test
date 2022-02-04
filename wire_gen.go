// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/chandraekap/fita-backend-test/graph"
	"github.com/chandraekap/fita-backend-test/internal/inventory"
	mock2 "github.com/chandraekap/fita-backend-test/internal/inventory/mock"
	"github.com/chandraekap/fita-backend-test/internal/sales"
	"github.com/chandraekap/fita-backend-test/internal/sales/mock"
	"github.com/google/wire"
)

// Injectors from wire.go:

func BuildResolver() *graph.Resolver {
	data := mock.InitData()
	v := data.Carts
	cartRepository := mock.NewCartRepository(v)
	mockData := mock2.InitData()
	v2 := mockData.Items
	itemRepository := mock2.NewItemRepository(v2)
	v3 := mockData.Stocks
	stockRepository := mock2.NewStockRepository(v3)
	v4 := data.Promos
	promotionRepository := mock.NewPromotionRepository(v4)
	promotionItemFactory := sales.NewPromotionItemFactory()
	checkoutService := sales.NewCheckoutService(cartRepository, promotionRepository, promotionItemFactory)
	resolver := graph.NewResolver(cartRepository, itemRepository, stockRepository, checkoutService)
	return resolver
}

// wire.go:

var salesSet = wire.NewSet(mock.InitData, wire.FieldsOf(new(*mock.Data), "Carts", "Promos"), mock.NewCartRepository, wire.Bind(new(sales.CartRepository), new(*mock.CartRepository)), mock.NewPromotionRepository, wire.Bind(new(sales.PromotionRepository), new(*mock.PromotionRepository)), sales.NewPromotionItemFactory, sales.NewCheckoutService)

var inventorySet = wire.NewSet(mock2.InitData, wire.FieldsOf(new(*mock2.Data), "Items", "Stocks"), mock2.NewItemRepository, wire.Bind(new(inventory.ItemRepository), new(*mock2.ItemRepository)), mock2.NewStockRepository, wire.Bind(new(inventory.StockRepository), new(*mock2.StockRepository)))