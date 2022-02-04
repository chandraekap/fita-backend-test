package test

import "github.com/chandraekap/fita-backend-test/internal/sales"

type TestContainer struct {
	CheckoutService sales.CheckoutService
}

func NewTestContainer(
	checkoutService sales.CheckoutService,
) *TestContainer {
	return &TestContainer{
		CheckoutService: checkoutService,
	}
}
