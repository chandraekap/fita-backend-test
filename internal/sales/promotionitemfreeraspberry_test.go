package sales_test

import (
	"context"
	"testing"

	"github.com/chandraekap/fita-backend-test/internal/sales"
	salestest "github.com/chandraekap/fita-backend-test/internal/sales/test"
	"github.com/google/go-cmp/cmp"
)

func TestPromotionItemFreeRaspberry(t *testing.T) {
	var testCases = map[string]struct {
		cartItem *sales.CartItem
		expected *sales.PromoItem
		err      error
	}{
		"success case": {
			cartItem: &sales.CartItem{
				SKU:   "43N23P",
				Name:  "MacBook Pro",
				Price: 5399.99,
				Qty:   1,
			},
			expected: &sales.PromoItem{
				SKU:   "234234",
				Price: 0,
				Qty:   1,
			},
			err: nil,
		},
		"negative case qty greater than 1": {
			cartItem: &sales.CartItem{
				SKU:   "43N23P",
				Name:  "MacBook Pro",
				Price: 5399.99,
				Qty:   2,
			},
			expected: &sales.PromoItem{
				SKU:   "234234",
				Price: 0,
				Qty:   2,
			},
			err: nil,
		},
	}

	container := salestest.BuildTestContainer()

	for caseName, tc := range testCases {
		actual, err := container.PromotionItemFreeRaspberry.Get(context.TODO(), tc.cartItem)
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
