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
	Type        string
	Amount      float64
	Method      string
	Currency    null.String
	BankName    null.String
	QrCode      null.String
	PaymentLink null.String
	ExpiredAt   null.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
