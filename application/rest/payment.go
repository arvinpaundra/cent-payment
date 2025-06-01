package rest

import (
	"net/http"

	"github.com/arvinpaundra/cent/payment/core/format"
	"github.com/arvinpaundra/cent/payment/domain/donation/dto/request"
	"github.com/arvinpaundra/cent/payment/domain/donation/service"
	"github.com/arvinpaundra/cent/payment/infrastructure/donation"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (cont Controller) CreateDonation(c *gin.Context) {
	var payload request.CreateDonation

	_ = c.ShouldBindJSON(&payload)

	verrs := cont.validator.Validate(payload)
	if verrs != nil {
		c.JSON(http.StatusBadRequest, format.BadRequest("invalid request body", verrs))
		return
	}

	handler := service.NewCreateDonationHandler(
		donation.NewPaymentWriterRepository(cont.db),
		donation.NewMidtrans(viper.GetString("MIDTRANS_SERVER_KEY"), viper.GetString("MIDTRANS_MODE")),
	)

	res, err := handler.Handle(c, payload)
	if err != nil {
		switch err {
		default:
			c.JSON(http.StatusInternalServerError, format.InternalServerError())
			return
		}
	}

	c.JSON(http.StatusCreated, format.SuccessCreated("success create donation", gin.H{
		"payment_url": res,
	}))
}
