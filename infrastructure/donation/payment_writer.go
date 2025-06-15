package donation

import (
	"context"
	"database/sql"

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

func (r PaymentWriterRepository) Save(ctx context.Context, payment *entity.Payment) error {
	if payment.IsNew() {
		return r.insert(ctx, payment)
	} else if payment.IsMarkedToBeUpdated() {
		return r.update(ctx, payment)
	}

	return nil
}

func (r PaymentWriterRepository) insert(ctx context.Context, payment *entity.Payment) error {
	paymentModel := model.Payment{
		UserId:  payment.UserId,
		Code:    payment.Code,
		Source:  payment.Source.String(),
		Status:  payment.Status.String(),
		Method:  payment.Method.String(),
		Purpose: payment.Purpose.String(),
		Amount:  payment.Amount,
	}

	err := r.db.WithContext(ctx).
		Model(&model.Payment{}).
		Create(&paymentModel).
		Error

	if err != nil {
		return err
	}

	payment.ID = paymentModel.ID

	if payment.HasDetail() {
		pd := payment.PaymentDetail

		var campaignId null.Int64

		if pd.CampaignId != nil {
			campaignId = null.Int64{NullInt64: sql.NullInt64{Int64: *pd.CampaignId, Valid: true}}
		}

		paymentDetailModel := model.PaymentDetail{
			PaymentId:  paymentModel.ID,
			ContentId:  pd.ContentId,
			Name:       pd.Name,
			Message:    pd.Message,
			Phone:      null.StringFromPtr(pd.Phone),
			Email:      null.StringFromPtr(pd.Email),
			CampaignId: campaignId,
		}

		err = r.db.WithContext(ctx).
			Model(&model.PaymentDetail{}).
			Create(&paymentDetailModel).
			Error

		if err != nil {
			return err
		}

		pd.ID = paymentDetailModel.ID
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
		QrCode:    null.StringFromPtr(payment.Qrcode),
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

	if payment.HasDetail() {
		pd := payment.PaymentDetail

		var campaignId null.Int64

		if pd.CampaignId != nil {
			campaignId = null.Int64{NullInt64: sql.NullInt64{Int64: *pd.CampaignId, Valid: true}}
		}

		paymentDetailModel := model.PaymentDetail{
			Name:       pd.Name,
			Message:    pd.Message,
			Phone:      null.StringFromPtr(pd.Phone),
			Email:      null.StringFromPtr(pd.Email),
			CampaignId: campaignId,
		}

		err = r.db.WithContext(ctx).
			Model(&model.PaymentDetail{}).
			Where("id = ?", pd.ID).
			Updates(&paymentDetailModel).
			Error

		if err != nil {
			return err
		}
	}

	return nil
}
