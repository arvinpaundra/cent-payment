package constant

import "errors"

var (
	ErrPaymentNotFound     = errors.New("payment not found")
	ErrPaymentAlreadyPaid  = errors.New("payment already paid")
	ErrInvalidSignatureKey = errors.New("invalid signature key")
	ErrInequalPaidAmount   = errors.New(("amount initiated isn't equal with paid amount"))
)
