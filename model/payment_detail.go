package model

import (
	"time"

	"github.com/guregu/null/v6"
)

type PaymentDetail struct {
	ID         int64
	PaymentId  int64
	ContentId  int64
	Name       string
	Message    string
	CampaignId null.Int64
	Phone      null.String
	Email      null.String
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
