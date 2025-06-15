package model

import (
	"time"

	"github.com/guregu/null/v6"
)

type Payment struct {
	ID          int64
	UserId      int64
	Code        string
	Source      string
	Status      string
	Method      string
	Purpose     string
	Amount      float64
	Reference   null.String
	Currency    null.String
	QrCode      null.String
	PaymentLink null.String
	ExpiredAt   null.Time
	PaidAt      null.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time

	PaymentDetail *PaymentDetail
}
