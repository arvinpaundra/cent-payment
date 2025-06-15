package webhook

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/webhook/entity"
	"github.com/arvinpaundra/cent/payment/model"
	"github.com/guregu/null/v6"
	"gorm.io/gorm"
)

type PaymentWriterRepository struct {
	db *gorm.DB
}

func NewPaymentWriterRepository(db *gorm.DB) PaymentWriterRepository {
	return PaymentWriterRepository{db: db}
}

func (r PaymentWriterRepository) Save(ctx context.Context, payment *entity.Payment) error {
	if payment.IsMarkedToBeUpdated() {
		return r.update(ctx, payment)
	}

	return nil
}

func (r PaymentWriterRepository) update(ctx context.Context, payment *entity.Payment) error {
	paymentModel := model.Payment{
		UserId:    payment.UserId,
		Status:    payment.Status.String(),
		Method:    payment.Method.String(),
		Amount:    payment.Amount,
		Reference: null.StringFromPtr(payment.Reference),
		Currency:  null.StringFromPtr(payment.Currency),
		PaidAt:    null.TimeFromPtr(payment.PaidAt),
	}

	err := r.db.WithContext(ctx).
		Model(&model.Payment{}).
		Where("id = ?", payment.ID).
		Updates(&paymentModel).
		Error

	if err != nil {
		return err
	}

	return nil
}
