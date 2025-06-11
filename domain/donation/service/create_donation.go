package service

import (
	"context"

	donationcmd "github.com/arvinpaundra/cent/payment/application/command/donation"
	"github.com/arvinpaundra/cent/payment/domain/donation/constant"
	"github.com/arvinpaundra/cent/payment/domain/donation/entity"
	"github.com/arvinpaundra/cent/payment/domain/donation/repository"
)

type CreateDonationHandler struct {
	paymentWriter       repository.PaymentWriter
	paymentGateway      repository.PaymentGateway
	unitOfWork          repository.UnitOfWork
	userClientMapper    repository.UserClientMapper
	contentClientMapper repository.ContentClientMapper
}

func NewCreateDonationHandler(
	paymentWriter repository.PaymentWriter,
	paymentGateway repository.PaymentGateway,
	unitOfWork repository.UnitOfWork,
	userClientMapper repository.UserClientMapper,
	contentClientMapper repository.ContentClientMapper,
) CreateDonationHandler {
	return CreateDonationHandler{
		paymentWriter:       paymentWriter,
		paymentGateway:      paymentGateway,
		unitOfWork:          unitOfWork,
		userClientMapper:    userClientMapper,
		contentClientMapper: contentClientMapper,
	}
}

func (s CreateDonationHandler) Handle(ctx context.Context, command donationcmd.CreateDonation) (*string, error) {
	// find user by slug
	user, err := s.userClientMapper.FindUserDetail(ctx, command.UserSlug)
	if err != nil {
		return nil, err
	}

	payment := entity.Payment{
		UserId: user.ID,
		Source: constant.PaymentSourceMidtrans,
		Type:   constant.PaymentTypeDonation,
		Status: constant.PaymentStatusPending,
		Method: constant.PaymentMethodNone,
		Amount: command.Amount,
	}

	err = payment.GenerateCode()
	if err != nil {
		return nil, err
	}

	content, err := s.contentClientMapper.FindActiveContent(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	var campaignId *int64

	if content.CampaignId != 0 {
		campaignId = &content.CampaignId
	}

	paymentDetail := &entity.PaymentDetail{
		ContentId:  content.ID,
		Name:       command.Name,
		Message:    command.Message,
		CampaignId: campaignId,
		Email:      command.Email,
		Phone:      command.Phone,
	}

	payment.SetPaymentDetail(paymentDetail)

	tx, err := s.unitOfWork.Begin()
	if err != nil {
		return nil, err
	}

	err = tx.PaymentWriter().Save(ctx, &payment)
	if err != nil {
		if uowErr := tx.Rollback(); uowErr != nil {
			return nil, uowErr
		}

		return nil, err
	}

	paymentGateway := entity.PaymentGateway{
		Amount: payment.Amount,
		Code:   payment.Code,
	}

	// create payment through payment gateway
	paymentLink, err := s.paymentGateway.Pay(ctx, &paymentGateway)
	if err != nil {
		if uowErr := tx.Rollback(); uowErr != nil {
			return nil, uowErr
		}

		return nil, err
	}

	payment.SetPaymentLink(paymentLink)

	payment.MarkToBeUpdated()

	err = tx.PaymentWriter().Save(ctx, &payment)
	if err != nil {
		if uowErr := tx.Rollback(); uowErr != nil {
			return nil, uowErr
		}

		return nil, err
	}

	if uowErr := tx.Commit(); uowErr != nil {
		return nil, uowErr
	}

	return &paymentLink, nil
}
