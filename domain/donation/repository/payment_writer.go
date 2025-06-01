package repository

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/entity"
)

type PaymentWriter interface {
	Save(ctx context.Context, payment entity.Payment) error
}
