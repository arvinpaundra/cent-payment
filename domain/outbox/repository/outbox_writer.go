package repository

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/outbox/entity"
)

type OutboxWriter interface {
	Save(ctx context.Context, outbox *entity.Outbox) error
}
