package handler

import (
	"net/http"

	webhookcmd "github.com/arvinpaundra/cent/payment/application/command/webhook"
	"github.com/arvinpaundra/cent/payment/core/format"
	"github.com/arvinpaundra/cent/payment/domain/webhook/constant"
	"github.com/arvinpaundra/cent/payment/domain/webhook/service"
	webhookinfra "github.com/arvinpaundra/cent/payment/infrastructure/webhook"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (h Handler) MidtransUpdatePayment(c *gin.Context) {
	var command webhookcmd.MidtransUpdateWebhook

	_ = c.ShouldBindJSON(&command)

	svc := service.NewMidtransUpdatePayment(
		webhookinfra.NewPaymentReaderRepository(h.db),
		webhookinfra.NewPaymentWriterRepository(h.db),
		webhookinfra.NewOutboxWriterRepository(h.db),
		webhookinfra.NewUnitOfWork(h.db),
		webhookinfra.NewUserClientMapper(
			h.grpcClient.UserClient(),
			viper.GetString("USER_CLIENT_API_KEY"),
		),
	)

	err := svc.Exec(c, command)
	if err != nil {
		switch err {
		case constant.ErrPaymentNotFound, constant.ErrPaymentAlreadyPaid, constant.ErrInequalPaidAmount:
			c.JSON(http.StatusUnprocessableEntity, format.UnprocessableEntity(err.Error()))
			return
		case constant.ErrInvalidSignatureKey:
			c.JSON(http.StatusForbidden, format.Forbidden(err.Error()))
			return
		default:
			c.JSON(http.StatusInternalServerError, format.InternalServerError())
			return
		}
	}

	c.JSON(http.StatusOK, format.SuccessOK("webhook proceeded successfully", nil))
}
