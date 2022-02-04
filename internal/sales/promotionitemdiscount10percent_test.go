package sales_test

import (
	"context"
	"testing"

	"github.com/chandraekap/fita-backend-test/internal/sales"
	salestest "github.com/chandraekap/fita-backend-test/internal/sales/test"
	"github.com/google/go-cmp/cmp"
)

func TestPromotionItemDiscount10Percent(t *testing.T) {
	var testCases = map[string]struct {
		cartItem *sales.CartItem
		expected *sales.PromoItem
		err      error
	}{
		"success case": {
			cartItem: &sales.CartItem{
				SKU:   "A304SD",
				Name:  "Alexa Speaker",
				Price: 109.50,
				Qty:   3,
			},
			expected: &sales.PromoItem{
				SKU:   "A304SD",
				Price: 98.55,
				Qty:   3,
			},
			err: nil,
		},
		"negative case qty greater than 3": {
			cartItem: &sales.CartItem{
				SKU:   "A304SD",
				Name:  "Alexa Speaker",
				Price: 109.50,
				Qty:   10,
			},
			expected: &sales.PromoItem{
				SKU:   "A304SD",
				Price: 98.55,
				Qty:   10,
			},
			err: nil,
		},
	}

	container := salestest.BuildTestContainer()

	for caseName, tc := range testCases {
		actual, err := container.PromotionItemDiscount10Percent.Get(context.TODO(), tc.cartItem)
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
