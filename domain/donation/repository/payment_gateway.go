package repository

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/external"
)

type PaymentGateway interface {
	Pay(ctx context.Context, pg *external.PaymentGatewayRequest) (*external.PaymentGatewayResponse, error)
}
