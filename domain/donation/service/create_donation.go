package service

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/constant"
	"github.com/arvinpaundra/cent/payment/domain/donation/dto/request"
	"github.com/arvinpaundra/cent/payment/domain/donation/entity"
	"github.com/arvinpaundra/cent/payment/domain/donation/repository"
)

type CreateDonationHandler struct {
	paymentWriter  repository.PaymentWriter
	paymentGateway repository.PaymentGateway
}

func NewCreateDonationHandler(
	paymentWriter repository.PaymentWriter,
	paymentGateway repository.PaymentGateway,
) CreateDonationHandler {
	return CreateDonationHandler{
		paymentWriter:  paymentWriter,
		paymentGateway: paymentGateway,
	}
}

func (s CreateDonationHandler) Handle(ctx context.Context, payload request.CreateDonation) (string, error) {
	payment := entity.Payment{
		UserId: payload.UserId,
		Source: constant.PaymentSourceMidtrans,
		Type:   constant.PaymentTypeDonation,
		Status: constant.PaymentStatusPending,
		Amount: payload.Amount,
	}

	err := payment.GenerateCode()
	if err != nil {
		return "", err
	}

	paymentDetail := &entity.PaymentDetail{
		Name:  payload.Name,
		Email: &payload.Email,
		Phone: &payload.Phone,
	}

	payment.SetPaymentDetail(paymentDetail)

	err = s.paymentWriter.Save(ctx, payment)
	if err != nil {
		return "", err
	}

	paymentGateway := entity.PaymentGateway{
		Amount: payment.Amount,
		Code:   payment.Code,
	}

	// create payment through payment gateway
	paymentUrl, err := s.paymentGateway.Pay(ctx, paymentGateway)
	if err != nil {
		return "", err
	}

	return paymentUrl, nil
}
