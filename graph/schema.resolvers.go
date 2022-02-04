package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/chandraekap/fita-backend-test/graph/generated"
	"github.com/chandraekap/fita-backend-test/graph/model"
	"github.com/chandraekap/fita-backend-test/internal/sales"
)

func (r *mutationResolver) AddCart(ctx context.Context, input model.CartAddRequest) (*model.Cart, error) {
	item, err := r.itemRepository.FindBySKU(ctx, input.Sku)
	if err != nil {
		return nil, err
	}

	if item.InventoryQty < input.Qty {
		return nil, errors.New("Item stock less than input quantity.")
	}

	activeCart, err := r.cartRepository.FindByClientID(ctx, input.ClientID)
	if err != nil && err != sales.ErrCartNotFound {
		return nil, err
	}
	if activeCart == nil {
		activeCart = &sales.Cart{
			ClientID: input.ClientID,
			Items: []*sales.CartItem{
				{
					SKU:   item.SKU,
					Name:  item.Name,
					Qty:   input.Qty,
					Price: item.Price,
				},
			},
		}
		err = r.cartRepository.Insert(ctx, activeCart)
		if err != nil {
			return nil, err
		}

	} else {
		for _, cartItem := range activeCart.Items {
			if cartItem.SKU == item.SKU {
				cartItem.Qty += input.Qty
			}
		}

		err = r.cartRepository.Update(ctx, activeCart)
		if err != nil {
			return nil, err
		}
	}

	cartItemsView := []*model.CartItem{}
	for _, cartItem := range activeCart.Items {
		cartItemsView = append(cartItemsView, &model.CartItem{
			Sku:           cartItem.SKU,
			Name:          cartItem.Name,
			Qty:           cartItem.Qty,
			Price:         cartItem.Price,
			TotalDiscount: cartItem.TotalDiscount,
		})
	}

	return &model.Cart{
		ID:       &activeCart.ID,
		ClientID: activeCart.ClientID,
		Items:    cartItemsView,
	}, nil
}

func (r *mutationResolver) Checkout(ctx context.Context, input model.CheckoutRequest) (*model.CheckoutSummary, error) {
	activeCart, err := r.cartRepository.FindByClientID(ctx, input.ClientID)
	if err != nil && err != sales.ErrCartNotFound {
		return nil, err
	}

	if activeCart != nil {
		for _, cartItem := range activeCart.Items {
			item, err := r.itemRepository.FindBySKU(ctx, cartItem.SKU)
			if err != nil {
				return nil, err
			}
			if item.InventoryQty < cartItem.Qty {
				return nil, errors.New("Item stock less than input quantity.")
			}
		}
	}

	checkoutSummary, err := r.checkoutService.Checkout(ctx, input.ClientID)
	if err != nil {
		return nil, err
	}

	items := []*model.CartItem{}
	for _, item := range checkoutSummary.Items {
		items = append(items, &model.CartItem{
			Sku:           item.SKU,
			Name:          item.Name,
			Qty:           item.Qty,
			Price:         item.Price,
			TotalDiscount: item.TotalDiscount,
		})
	}

	return &model.CheckoutSummary{
		Items:          items,
		DiscountAmount: checkoutSummary.DiscountAmount,
		TotalAmount:    checkoutSummary.TotalAmount,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type queryResolver struct{ *Resolver }
