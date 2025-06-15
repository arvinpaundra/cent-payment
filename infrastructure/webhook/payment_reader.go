package webhook

import (
	"context"
	"errors"

	"github.com/arvinpaundra/cent/payment/domain/webhook/constant"
	"github.com/arvinpaundra/cent/payment/domain/webhook/entity"
	"github.com/arvinpaundra/cent/payment/model"
	"gorm.io/gorm"
)

type PaymentReaderRepository struct {
	db *gorm.DB
}

func NewPaymentReaderRepository(db *gorm.DB) PaymentReaderRepository {
	return PaymentReaderRepository{db: db}
}

func (r PaymentReaderRepository) FindByCode(ctx context.Context, code string) (*entity.Payment, error) {
	var paymentModel model.Payment

	err := r.db.WithContext(ctx).
		Model(&model.Payment{}).
		Preload("PaymentDetail").
		Where("code = ?", code).
		First(&paymentModel).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrPaymentNotFound
		}

		return nil, err
	}

	payment := entity.Payment{
		ID:     paymentModel.ID,
		UserId: paymentModel.UserId,
		Amount: paymentModel.Amount,
		Method: constant.PaymentMethod(paymentModel.Method),
		Status: constant.PaymentStatus(paymentModel.Status),
		PaidAt: paymentModel.PaidAt.Ptr(),
	}

	if paymentModel.PaymentDetail != nil {
		paymentDetailModel := paymentModel.PaymentDetail

		paymentDetail := &entity.PaymentDetail{
			ID:        paymentDetailModel.ID,
			PaymentId: paymentModel.ID,
			ContentId: paymentDetailModel.ContentId,
			Message:   paymentDetailModel.Message,
			Name:      paymentDetailModel.Name,
		}

		payment.PaymentDetail = paymentDetail
	}

	return &payment, nil
}
