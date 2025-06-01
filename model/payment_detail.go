package model

import (
	"time"

	"github.com/guregu/null/v6"
)

type PaymentDetail struct {
	ID        int64
	PaymentId int64
	Name      string
	Phone     null.String
	Email     null.String
	CreatedAt time.Time
	UpdatedAt time.Time
}
