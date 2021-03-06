// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Cart struct {
	ID       *int        `json:"id"`
	ClientID int         `json:"clientID"`
	Items    []*CartItem `json:"items"`
}

type CartAddRequest struct {
	ClientID int    `json:"clientID"`
	Sku      string `json:"sku"`
	Qty      int    `json:"qty"`
}

type CartItem struct {
	Sku   string  `json:"sku"`
	Name  string  `json:"name"`
	Qty   int     `json:"qty"`
	Price float64 `json:"price"`
}

type CheckoutRequest struct {
	ClientID int `json:"clientID"`
}

type CheckoutSummary struct {
	Items          []*CartItem `json:"items"`
	DiscountAmount float64     `json:"discountAmount"`
	TotalAmount    float64     `json:"totalAmount"`
}

type Item struct {
	Sku          string  `json:"sku"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	InventoryQty int     `json:"inventoryQty"`
}
