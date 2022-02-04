package mock

import "github.com/chandraekap/fita-backend-test/internal/inventory"

type Data struct {
	Items  []*inventory.Item
	Stocks []*inventory.Stock
}

func InitData() *Data {
	return &Data{
		Items: []*inventory.Item{
			{
				SKU:          "120P90",
				Name:         "Google Home",
				Price:        49.99,
				InventoryQty: 10,
			},
			{
				SKU:          "43N23P",
				Name:         "MacBook Pro",
				Price:        5399.99,
				InventoryQty: 5,
			},
			{
				SKU:          "A304SD",
				Name:         "Alexa Speaker",
				Price:        109.50,
				InventoryQty: 10,
			},
			{
				SKU:          "234234",
				Name:         "Raspberry Pi B",
				Price:        30.00,
				InventoryQty: 2,
			},
		},
		Stocks: []*inventory.Stock{
			{
				ID:     1,
				SKU:    "120P90",
				Reason: "addition",
				Qty:    10,
			},
			{
				ID:     2,
				SKU:    "43N23P",
				Reason: "addition",
				Qty:    5,
			},
			{
				ID:     3,
				SKU:    "A304SD",
				Reason: "addition",
				Qty:    10,
			},
			{
				ID:     4,
				SKU:    "234234",
				Reason: "addition",
				Qty:    2,
			},
		},
	}
}
