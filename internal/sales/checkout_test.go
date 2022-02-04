package sales_test

import (
	"context"
	"testing"

	"github.com/chandraekap/fita-backend-test/internal/sales"
	salestest "github.com/chandraekap/fita-backend-test/internal/sales/test"
	"github.com/google/go-cmp/cmp"
)

func TestCheckout(t *testing.T) {
	var testCases = map[string]struct {
		clientID int
		expected *sales.CheckoutSummary
		err      error
	}{
		"success case macbook": {
			clientID: 99999,
			expected: &sales.CheckoutSummary{
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
						Price: 0,
						Qty:   1,
					},
				},
				TotalAmount: 5399.99,
			},
			err: nil,
		},
		"negative case macbook with promo in buying": {
			clientID: 99998,
			expected: &sales.CheckoutSummary{
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
						Qty:   2,
					},
					{
						SKU:   "234234",
						Name:  "Raspberry Pi B",
						Price: 0,
						Qty:   1,
					},
				},
				TotalAmount: 5459.99,
			},
			err: nil,
		},
		"success case google home": {
			clientID: 88888,
			expected: &sales.CheckoutSummary{
				Items: []*sales.CartItem{
					{
						SKU:   "120P90",
						Name:  "Google Home",
						Price: 49.99,
						Qty:   2,
					},
					{
						SKU:   "120P90",
						Name:  "Google Home",
						Price: 0,
						Qty:   1,
					},
				},
				TotalAmount: 99.98,
			},
			err: nil,
		},
		"success case alexa speaker": {
			clientID: 77777,
			expected: &sales.CheckoutSummary{
				Items: []*sales.CartItem{
					{
						SKU:   "A304SD",
						Name:  "Alexa Speaker",
						Price: 98.55,
						Qty:   3,
					},
				},
				TotalAmount: 295.65,
			},
			err: nil,
		},
	}

	container := salestest.BuildTestContainer()

	for caseName, tc := range testCases {
		actual, err := container.CheckoutService.Checkout(context.TODO(), tc.clientID)
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
