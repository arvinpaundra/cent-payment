package repository

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/webhook/data"
)

type UserClientMapper interface {
	UpdateUserBalance(ctx context.Context, payload *data.UpdateBalanceUserRequest) error
}
