package service

import "context"

type MidtransUpdatePaymentHandler struct {
}

func NewMidtransUpdatePaymentHandler() MidtransUpdatePaymentHandler {
	return MidtransUpdatePaymentHandler{}
}

func (s MidtransUpdatePaymentHandler) Handle(ctx context.Context) error {
	return nil
}
