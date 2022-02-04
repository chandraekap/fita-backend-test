package sales

import "context"

type PromotionItemFactory interface {
	GetService(ctx context.Context, ruleType RuleType) PromotionItemService
}

type promotionItemFactory struct {
}

func NewPromotionItemFactory() PromotionItemFactory {
	return &promotionItemFactory{}
}

func (factory *promotionItemFactory) GetService(ctx context.Context, ruleType RuleType) PromotionItemService {
	if ruleType == RULE_FREE_RASPBERRY {
		return NewPromotionItemFreeRaspberry()
	} else if ruleType == RULE_BUY_2_FREE_1 {
		return NewPromotionItemBuy2Free1()
	} else if ruleType == RULE_DISCOUNT_10PERCENT {
		return NewPromotionItemDiscount10Percent()
	}

	return nil
}
