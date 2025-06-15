package repository

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/webhook/entity"
)

type PaymentReader interface {
	FindByCode(ctx context.Context, code string) (*entity.Payment, error)
}
