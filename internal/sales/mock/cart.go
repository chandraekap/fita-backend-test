package mock

import (
	"context"
	"errors"

	"github.com/chandraekap/fita-backend-test/internal/sales"
)

type CartRepository struct {
	data []*sales.Cart
}

func NewCartRepository(
	data []*sales.Cart,
) *CartRepository {
	return &CartRepository{
		data: data,
	}
}

func (repo *CartRepository) FindByClientID(ctx context.Context, clientID int) (*sales.Cart, error) {
	for _, cartData := range repo.data {
		if cartData.ClientID == clientID {
			return cartData, nil
		}
	}

	return nil, errors.New("Cart not found.")
}

func (repo *CartRepository) Insert(ctx context.Context, cart *sales.Cart) error {
	for _, cartData := range repo.data {
		if cartData.ClientID == cart.ClientID {
			repo.data = append(repo.data, cart)
			return nil
		}
	}

	return errors.New("Failed to insert Cart.")
}

func (repo *CartRepository) Update(ctx context.Context, cart *sales.Cart) error {
	for i, cartData := range repo.data {
		if cartData.ClientID == cart.ClientID {
			repo.data[i] = cart
			return nil
		}
	}

	return errors.New("Failed to insert sales.")
}

func (repo *CartRepository) Delete(ctx context.Context, cart *sales.Cart) error {
	for i, cartData := range repo.data {
		if cartData.ClientID == cart.ClientID {
			repo.data = append(repo.data[:i], repo.data[i+1:]...)
			return nil
		}
	}

	return errors.New("Failed to insert Cart.")
}
