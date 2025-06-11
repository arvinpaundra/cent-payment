package repository

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/data"
)

type UserClientMapper interface {
	FindUserDetail(ctx context.Context, slug string) (*data.UserResponse, error)
}
