package rest

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

func (cont Controller) MidtransUpdatePayment(c *gin.Context) {
	var command webhookcmd.MidtransUpdateWebhook

	_ = c.ShouldBindJSON(&command)

	handler := service.NewMidtransUpdatePaymentHandler(
		webhookinfra.NewPaymentReaderRepository(cont.db),
		webhookinfra.NewPaymentWriterRepository(cont.db),
		webhookinfra.NewOutboxWriterRepository(cont.db),
		webhookinfra.NewUnitOfWork(cont.db),
		webhookinfra.NewUserClientMapper(
			cont.grpcClient.UserClient(),
			viper.GetString("USER_CLIENT_API_KEY"),
		),
	)

	err := handler.Handle(c, command)
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
