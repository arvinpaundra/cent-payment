package service

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	webhookcmd "github.com/arvinpaundra/cent/payment/application/command/webhook"
	"github.com/arvinpaundra/cent/payment/core/util"
	"github.com/arvinpaundra/cent/payment/domain/webhook/constant"
	"github.com/arvinpaundra/cent/payment/domain/webhook/entity"
	"github.com/arvinpaundra/cent/payment/domain/webhook/external"
	"github.com/arvinpaundra/cent/payment/domain/webhook/repository"
	"github.com/spf13/viper"
)

type MidtransUpdatePayment struct {
	paymentReader    repository.PaymentReader
	paymentWriter    repository.PaymentWriter
	outboxWriter     repository.OutboxWriter
	uow              repository.UnitOfWork
	userClientMapper repository.UserClientMapper
}

func NewMidtransUpdatePayment(
	paymentReader repository.PaymentReader,
	paymentWriter repository.PaymentWriter,
	outboxWriter repository.OutboxWriter,
	uow repository.UnitOfWork,
	userClientMapper repository.UserClientMapper,
) MidtransUpdatePayment {
	return MidtransUpdatePayment{
		paymentReader:    paymentReader,
		paymentWriter:    paymentWriter,
		outboxWriter:     outboxWriter,
		uow:              uow,
		userClientMapper: userClientMapper,
	}
}

func (s MidtransUpdatePayment) Exec(ctx context.Context, command webhookcmd.MidtransUpdateWebhook) error {
	payment, err := s.paymentReader.FindByCode(ctx, command.OrderId)
	if err != nil {
		return err
	}

	err = s.validate(payment, command)
	if err != nil {
		return err
	}

	paymentStatus := s.paymentStatus(command.TransactionStatus)

	payment.SetStatus(paymentStatus)
	payment.SetReference(command.TransactionId)

	paymentMethod := s.paymentMethod(command.PaymentType)

	payment.SetMethod(paymentMethod)

	if payment.IsPaid() {
		paidAt, _ := util.StringToTime(time.DateTime, command.SettlementTime)

		payment.SetPaidAt(paidAt)
		payment.SetCurrency(command.Currency)
	}

	payment.MarkToBeUpdated()

	tx, err := s.uow.Begin()
	if err != nil {
		return err
	}

	err = tx.PaymentWriter().Save(ctx, payment)
	if err != nil {
		if uowErr := tx.Rollback(); uowErr != nil {
			return uowErr
		}

		return err
	}

	if payment.IsPaid() {
		user, err := s.userClientMapper.FindUserDetail(ctx, payment.UserId)
		if err != nil {
			if uowErr := tx.Rollback(); uowErr != nil {
				return uowErr
			}

			return err
		}

		outboxPayload := struct {
			UserId  int64   `json:"user_id"`
			UserKey string  `json:"user_key"`
			Amount  float64 `json:"amount"`
			Sender  string  `json:"sender"`
			Message string  `json:"message"`
		}{
			UserId:  payment.UserId,
			UserKey: user.Key,
			Amount:  payment.Amount,
			Sender:  payment.PaymentDetail.Name,
			Message: payment.PaymentDetail.Message,
		}

		payloadBytes, err := json.Marshal(outboxPayload)
		if err != nil {
			if uowErr := tx.Rollback(); uowErr != nil {
				return uowErr
			}

			return err
		}

		outbox := &entity.Outbox{
			Status:  constant.OutboxStatusPending,
			Event:   constant.OutboxEventDonationPaid,
			Payload: payloadBytes,
		}

		err = tx.OutboxWriter().Save(ctx, outbox)
		if err != nil {
			if uowErr := tx.Rollback(); uowErr != nil {
				return uowErr
			}

			return err
		}

		updateBalancePayload := external.UpdateBalanceUserRequest{
			UserId: payment.UserId,
			Amount: payment.Amount,
		}

		err = s.userClientMapper.UpdateUserBalance(ctx, &updateBalancePayload)
		if err != nil {
			if uowErr := tx.Rollback(); uowErr != nil {
				return uowErr
			}

			return err
		}
	}

	if uowErr := tx.Commit(); uowErr != nil {
		return uowErr
	}

	return nil
}

func (s MidtransUpdatePayment) paymentStatus(status string) constant.PaymentStatus {
	switch status {
	case "settlement":
		return constant.PaymentStatusPaid
	case "expiry":
		return constant.PaymentStatusExpired
	case "pending":
		return constant.PaymentStatusPending
	default:
		return constant.PaymentStatusFailed
	}
}

func (s MidtransUpdatePayment) paymentMethod(method string) constant.PaymentMethod {
	switch method {
	case "gopay":
		return constant.PaymentMethodGopay
	case "shopeepay":
		return constant.PaymentMethodShopeepay
	case "qris":
		return constant.PaymentMethodQris
	default:
		return constant.PaymentMethodOthers
	}
}

func (s MidtransUpdatePayment) validate(payment *entity.Payment, command webhookcmd.MidtransUpdateWebhook) error {
	payload := fmt.Sprintf("%s+%s+%s+%s", command.OrderId, command.StatusCode, command.GrossAmount, viper.GetString("MIDTRANS_SERVER_KEY"))

	hash := sha512.Sum512([]byte(payload))
	signature := hex.EncodeToString(hash[:])

	if signature != command.SignatureKey {
		return constant.ErrInvalidSignatureKey
	}

	if payment.IsPaid() {
		return constant.ErrPaymentAlreadyPaid
	} else if command.TransactionStatus == "pending" {
		return nil
	}

	amount, _ := strconv.ParseFloat(command.GrossAmount, 64)

	if payment.Amount != amount {
		return constant.ErrInequalPaidAmount
	}

	return nil
}
