package mock

import "github.com/chandraekap/fita-backend-test/internal/sales"

type Data struct {
	Carts  []*sales.Cart
	Promos []*sales.Promotion
}

func InitData() *Data {
	return &Data{
		Carts: []*sales.Cart{
			{
				ID:       1,
				ClientID: 99999,
				Items: []*sales.CartItem{
					{
						SKU:   "43N23P",
						Name:  "MacBook Pro",
						Price: 5399.99,
						Qty:   1,
					},
					{
						SKU:   "234234",
						Name:  "Raspberry Pi B",
						Price: 30.00,
						Qty:   1,
					},
				},
			},
			{
				ID:       2,
				ClientID: 88888,
				Items: []*sales.CartItem{
					{
						SKU:   "120P90",
						Name:  "Google Home",
						Price: 49.99,
						Qty:   3,
					},
				},
			},
			{
				ID:       3,
				ClientID: 77777,
				Items: []*sales.CartItem{
					{
						SKU:   "A304SD",
						Name:  "Alexa Speaker",
						Price: 109.50,
						Qty:   3,
					},
				},
			},
			{
				ID:       4,
				ClientID: 99998,
				Items: []*sales.CartItem{
					{
						SKU:   "43N23P",
						Name:  "MacBook Pro",
						Price: 5399.99,
						Qty:   1,
					},
					{
						SKU:   "234234",
						Name:  "Raspberry Pi B",
						Price: 30.00,
						Qty:   3,
					},
				},
			},
		},
		Promos: []*sales.Promotion{
			{
				ItemSKU:  "43N23P",
				RuleType: sales.RULE_FREE_RASPBERRY,
			},
			{
				ItemSKU:  "120P90",
				RuleType: sales.RULE_BUY_2_FREE_1,
			},
			{
				ItemSKU:  "A304SD",
				RuleType: sales.RULE_DISCOUNT_10PERCENT,
			},
		},
	}
}
