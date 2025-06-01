package entity

import (
	"time"

	"github.com/arvinpaundra/cent/payment/domain/donation/constant"
)

type PaymentGateway struct {
	Amount float64
	Code   string
}

type PaymentGatewayResult struct {
	PaymentCode       string
	Currency          string
	VaNumber          string
	Amount            float64
	PaymentGatewayRef string
	ExpiryTime        time.Time
	Status            constant.PaymentStatus
	BankName          constant.Bank
	PaymentType       constant.PaymentMethod
}
