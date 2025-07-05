package handler

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

func (h Handler) CreateDonation(c *gin.Context) {
	var command donationcmd.CreateDonation

	_ = c.ShouldBindJSON(&command)

	verrs := h.validator.Validate(command)
	if verrs != nil {
		c.JSON(http.StatusBadRequest, format.BadRequest("invalid request body", verrs))
		return
	}

	slug := c.Param("slug")

	command.UserSlug = slug

	svc := service.NewCreateDonation(
		donationinfra.NewPaymentWriterRepository(h.db),
		donationinfra.NewMidtrans(
			viper.GetString("MIDTRANS_SERVER_KEY"),
			viper.GetString("MIDTRANS_MODE"),
		),
		donationinfra.NewUnitOfWork(h.db),
		donationinfra.NewUserClientMapper(
			h.grpcClient.UserClient(),
			viper.GetString("USER_CLIENT_API_KEY"),
		),
		donationinfra.NewContentClientMapper(
			h.grpcClient.ContentClient(),
		),
	)

	res, err := svc.Exec(c, command)
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
