package entity

import (
	"time"

	"github.com/arvinpaundra/cent/payment/core/trait"
	"github.com/arvinpaundra/cent/payment/domain/webhook/constant"
)

type Payment struct {
	trait.Updateable

	ID          int64
	UserId      int64
	Code        string
	Source      constant.PaymentSource
	Status      constant.PaymentStatus
	Method      constant.PaymentMethod
	Purpose     constant.PaymentPurpose
	Amount      float64
	Reference   *string
	Currency    *string
	PaymentLink *string
	ExpiredAt   *time.Time
	PaidAt      *time.Time

	PaymentDetail *PaymentDetail
}

func (e *Payment) IsPaid() bool {
	return e.Status == constant.PaymentStatusPaid
}

func (e *Payment) IsPending() bool {
	return e.Status == constant.PaymentStatusPending
}

func (e *Payment) SetMethod(method constant.PaymentMethod) {
	e.Method = method
}

func (e *Payment) SetStatus(status constant.PaymentStatus) {
	e.Status = status
}

func (e *Payment) SetReference(ref string) {
	e.Reference = &ref
}

func (e *Payment) SetCurrency(currency string) {
	e.Currency = &currency
}

func (e *Payment) SetPaidAt(t time.Time) {
	e.PaidAt = &t
}

type PaymentDetail struct {
	ID         int64
	PaymentId  int64
	ContentId  int64
	Name       string
	Message    string
	CampaignId *int64
	Phone      *string
	Email      *string
}
