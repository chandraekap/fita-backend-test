package graph

import (
	"github.com/chandraekap/fita-backend-test/internal/inventory"
	"github.com/chandraekap/fita-backend-test/internal/sales"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	cartRepository  sales.CartRepository
	itemRepository  inventory.ItemRepository
	stockRepository inventory.StockRepository
	checkoutService sales.CheckoutService
}

func NewResolver(
	cartRepository sales.CartRepository,
	itemRepository inventory.ItemRepository,
	stockRepository inventory.StockRepository,
	checkoutService sales.CheckoutService,
) *Resolver {
	return &Resolver{
		cartRepository:  cartRepository,
		itemRepository:  itemRepository,
		stockRepository: stockRepository,
		checkoutService: checkoutService,
	}
}
