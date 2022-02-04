package sales

import (
	"context"
)

type Cart struct {
	ID       int
	ClientID int
	Items    []*CartItem
}

type CartItem struct {
	SKU   string
	Name  string
	Qty   int
	Price float64
}

type CartRepository interface {
	FindByClientID(ctx context.Context, clientID int) (*Cart, error)
	Insert(ctx context.Context, cart *Cart) error
	Update(ctx context.Context, cart *Cart) error
	Delete(ctx context.Context, cart *Cart) error
}
