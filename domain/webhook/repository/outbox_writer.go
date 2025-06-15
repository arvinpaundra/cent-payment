package repository

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/webhook/entity"
)

type OutboxWriter interface {
	Save(ctx context.Context, outbox *entity.Outbox) error
}
