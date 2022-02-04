package sales

import apperrors "github.com/chandraekap/fita-backend-test/internal/errors"

var (
	ErrPromotionItemNotFound *apperrors.NotFoundError = apperrors.NewNotFoundError("Promotion Item")
	ErrCartNotFound          *apperrors.NotFoundError = apperrors.NewNotFoundError("Cart")
)
