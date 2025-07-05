package outbox

import (
	"context"
	"errors"

	"github.com/arvinpaundra/cent/payment/core/poller"
	"github.com/arvinpaundra/cent/payment/domain/outbox/constant"
	"github.com/arvinpaundra/cent/payment/domain/outbox/service"
	outboxinfra "github.com/arvinpaundra/cent/payment/infrastructure/outbox"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type Outbox struct {
	db *gorm.DB
	nc *nats.Conn
}

func NewOutbox(db *gorm.DB, nc *nats.Conn) Outbox {
	return Outbox{
		db: db,
		nc: nc,
	}
}

func (o Outbox) OutboxProcessor() error {
	svc := service.NewOutboxProcessor(
		outboxinfra.NewOutboxReaderRepository(o.db),
		outboxinfra.NewOutboxWriterRepository(o.db),
		outboxinfra.NewUnitOfWork(o.db),
		outboxinfra.NewMessaging(o.nc),
	)

	err := svc.Exec(context.Background())
	if err != nil {
		if errors.Is(err, constant.ErrOutboxNotFound) {
			return poller.ErrNoData
		}

		return err
	}

	return nil
}
