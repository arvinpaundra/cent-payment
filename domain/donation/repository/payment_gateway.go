package repository

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/entity"
)

type PaymentGateway interface {
	Pay(ctx context.Context, pg *entity.PaymentGateway) (string, error)
}
