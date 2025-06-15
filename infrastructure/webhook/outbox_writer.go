package webhook

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/webhook/entity"
	"github.com/arvinpaundra/cent/payment/model"
	"gorm.io/gorm"
)

type OutboxWriterRepository struct {
	db *gorm.DB
}

func NewOutboxWriterRepository(db *gorm.DB) OutboxWriterRepository {
	return OutboxWriterRepository{db: db}
}

func (r OutboxWriterRepository) Save(ctx context.Context, outbox *entity.Outbox) error {
	if outbox.IsNew() {
		return r.insert(ctx, outbox)
	}

	return nil
}

func (r OutboxWriterRepository) insert(ctx context.Context, outbox *entity.Outbox) error {
	outboxModel := model.Outbox{
		Status:  outbox.Status.String(),
		Event:   outbox.Event.String(),
		Payload: outbox.Payload,
	}

	err := r.db.WithContext(ctx).
		Model(&model.Outbox{}).
		Create(&outboxModel).
		Error

	if err != nil {
		return err
	}

	outbox.ID = outboxModel.ID

	return nil
}
