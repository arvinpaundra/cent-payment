package repository

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/external"
)

type ContentClientMapper interface {
	FindActiveContent(ctx context.Context, userId int64) (*external.ActiveContentResponse, error)
}
