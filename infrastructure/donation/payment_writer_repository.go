package donation

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/entity"
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

func (r PaymentWriterRepository) Save(ctx context.Context, payment entity.Payment) error {
	if payment.IsNew() {
		return r.insert(ctx, payment)
	}

	return nil
}

func (r PaymentWriterRepository) insert(ctx context.Context, payment entity.Payment) error {
	tx := r.db.Begin()

	paymentModel := model.Payment{
		UserId:    payment.UserId,
		Code:      payment.Code,
		Source:    payment.Source.String(),
		Status:    payment.Status.String(),
		Type:      payment.Type.String(),
		Amount:    payment.Amount,
		BankName:  null.StringFromPtr(payment.BankName),
		ExpiredAt: null.TimeFromPtr(payment.ExpiredAt),
	}

	err := tx.WithContext(ctx).
		Model(&model.Payment{}).
		Create(&paymentModel).
		Error

	if err != nil {
		tx.Rollback()
		return err
	}

	if payment.HasDetail() {
		pd := payment.PaymentDetail

		paymentDetailModel := model.PaymentDetail{
			PaymentId: paymentModel.ID,
			Name:      pd.Name,
			Phone:     null.StringFromPtr(pd.Phone),
			Email:     null.StringFromPtr(pd.Email),
		}

		err = tx.WithContext(ctx).
			Model(&model.PaymentDetail{}).
			Create(&paymentDetailModel).
			Error

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}
