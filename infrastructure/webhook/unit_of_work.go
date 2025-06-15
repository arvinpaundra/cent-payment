package webhook

import (
	"github.com/arvinpaundra/cent/payment/domain/webhook/repository"
	"gorm.io/gorm"
)

type UnitOfWork struct {
	db *gorm.DB
}

func (r UnitOfWork) Begin() (repository.UnitOfWorkProcessor, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return UnitOfWorkProcessor{tx: tx}, nil
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	return UnitOfWork{db: db}
}

type UnitOfWorkProcessor struct {
	tx *gorm.DB
}

func (r UnitOfWorkProcessor) PaymentWriter() repository.PaymentWriter {
	return NewPaymentWriterRepository(r.tx)
}

func (r UnitOfWorkProcessor) OutboxWriter() repository.OutboxWriter {
	return NewOutboxWriterRepository(r.tx)
}

func (r UnitOfWorkProcessor) Rollback() error {
	return r.tx.Rollback().Error
}

func (r UnitOfWorkProcessor) Commit() error {
	return r.tx.Commit().Error
}
