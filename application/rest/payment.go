package rest

import (
	"fmt"
	"net/http"

	donationcmd "github.com/arvinpaundra/cent/payment/application/command/donation"
	"github.com/arvinpaundra/cent/payment/core/format"
	"github.com/arvinpaundra/cent/payment/domain/donation/service"
	donationinfra "github.com/arvinpaundra/cent/payment/infrastructure/donation"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (cont Controller) CreateDonation(c *gin.Context) {
	var command donationcmd.CreateDonation

	_ = c.ShouldBindJSON(&command)

	verrs := cont.validator.Validate(command)
	if verrs != nil {
		c.JSON(http.StatusBadRequest, format.BadRequest("invalid request body", verrs))
		return
	}

	slug := c.Param("slug")

	command.UserSlug = slug

	handler := service.NewCreateDonationHandler(
		donationinfra.NewPaymentWriterRepository(cont.db),
		donationinfra.NewMidtrans(
			viper.GetString("MIDTRANS_SERVER_KEY"),
			viper.GetString("MIDTRANS_MODE"),
		),
		donationinfra.NewUnitOfWork(cont.db),
		donationinfra.NewUserClientMapper(
			cont.grpcClient.UserClient(),
			viper.GetString("USER_CLIENT_API_KEY"),
		),
		donationinfra.NewContentClientMapper(
			cont.grpcClient.ContentClient(),
		),
	)

	res, err := handler.Handle(c, command)
	if err != nil {
		fmt.Println("error:", err.Error())
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
