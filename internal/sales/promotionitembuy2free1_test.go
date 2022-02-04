package sales_test

import (
	"context"
	"testing"

	"github.com/chandraekap/fita-backend-test/internal/sales"
	salestest "github.com/chandraekap/fita-backend-test/internal/sales/test"
	"github.com/google/go-cmp/cmp"
)

func TestPromotionBuy2Free1(t *testing.T) {
	var testCases = map[string]struct {
		cartItem *sales.CartItem
		expected *sales.PromoItem
		err      error
	}{
		"success case": {
			cartItem: &sales.CartItem{
				SKU:   "120P90",
				Name:  "Google Home",
				Price: 49.99,
				Qty:   3,
			},
			expected: &sales.PromoItem{
				SKU:   "120P90",
				Price: 0,
				Qty:   1,
			},
			err: nil,
		},
		"negative case qty multiply of 3": {
			cartItem: &sales.CartItem{
				SKU:   "120P90",
				Name:  "Google Home",
				Price: 49.99,
				Qty:   9,
			},
			expected: &sales.PromoItem{
				SKU:   "120P90",
				Price: 0,
				Qty:   3,
			},
			err: nil,
		},
		"negative case qty not multiply of 3": {
			cartItem: &sales.CartItem{
				SKU:   "120P90",
				Name:  "Google Home",
				Price: 49.99,
				Qty:   8,
			},
			expected: &sales.PromoItem{
				SKU:   "120P90",
				Price: 0,
				Qty:   2,
			},
			err: nil,
		},
		"negative case qty less than 3": {
			cartItem: &sales.CartItem{
				SKU:   "120P90",
				Name:  "Google Home",
				Price: 49.99,
				Qty:   1,
			},
			expected: nil,
			err:      sales.ErrPromotionItemNotFound,
		},
	}

	container := salestest.BuildTestContainer()

	for caseName, tc := range testCases {
		actual, err := container.PromotionItemBuy2Free1.Get(context.TODO(), tc.cartItem)
		if err != tc.err {
			t.Fatalf("got %v, want %v", err, tc.err)
		}
		diff1 := cmp.Diff(tc.expected, actual)
		if diff1 != "" {
			t.Fatal("[", caseName, "]: ", diff1)
		}

		diff := cmp.Equal(tc.expected, actual)
		if !diff {
			t.Fatalf("got %+v, want %+v", actual, tc.expected)
		}
	}
}
