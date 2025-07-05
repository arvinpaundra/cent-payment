package repository

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/webhook/external"
)

type UserClientMapper interface {
	FindUserDetail(ctx context.Context, userId int64) (*external.UserResponse, error)	
	UpdateUserBalance(ctx context.Context, payload *external.UpdateBalanceUserRequest) error
}
