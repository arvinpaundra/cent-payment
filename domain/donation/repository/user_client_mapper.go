package repository

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/external"
)

type UserClientMapper interface {
	FindUserBySlug(ctx context.Context, slug string) (*external.UserResponse, error)
}
