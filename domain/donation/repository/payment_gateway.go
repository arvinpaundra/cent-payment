package repository

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/data"
)

type PaymentGateway interface {
	Pay(ctx context.Context, pg *data.PaymentGatewayRequest) (*data.PaymentGatewayResponse, error)
}
