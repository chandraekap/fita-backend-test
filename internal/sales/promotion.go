package sales

import "context"

type RuleType string

const RULE_FREE_RASPBERRY RuleType = "rule_raspberry_free"
const RULE_BUY_2_FREE_1 RuleType = "rule_buy_2_free_1"
const RULE_DISCOUNT_10PERCENT RuleType = "rule_discount_10percent"

type Promotion struct {
	ItemSKU  string
	MinQty   int
	RuleType RuleType
}

type PromotionRepository interface {
	FindBySKU(ctx context.Context, sku string) (*Promotion, error)
}
