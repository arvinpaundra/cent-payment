package repository

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/data"
)

type ContentClientMapper interface {
	FindActiveContent(ctx context.Context, userId int64) (*data.ActiveContentResponse, error)
}
