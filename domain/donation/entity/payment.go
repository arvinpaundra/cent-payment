package entity

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/arvinpaundra/cent/payment/core/trait"
	"github.com/arvinpaundra/cent/payment/domain/donation/constant"
)

type Payment struct {
	trait.Updateable

	ID          int64
	UserId      int64
	Code        string
	Source      constant.PaymentSource
	Method      constant.PaymentMethod
	Status      constant.PaymentStatus
	Purpose     constant.PaymentPurpose
	Amount      float64
	Reference   *string
	Currency    *string
	Qrcode      *string
	PaymentLink *string
	ExpiredAt   *time.Time
	PaidAt      *time.Time

	PaymentDetail *PaymentDetail
}

func (e *Payment) IsNew() bool {
	return e.ID <= 0
}

func (e *Payment) HasDetail() bool {
	return e.PaymentDetail != nil
}

func (e *Payment) GenerateCode() error {
	b := make([]byte, 3)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}

	randomHex := hex.EncodeToString(b)
	rndmHexUpper := strings.ToUpper(randomHex)

	e.Code = fmt.Sprintf("DON-%s", rndmHexUpper)

	return nil
}

func (e *Payment) SetExpiredAt(t time.Time) {
	e.ExpiredAt = &t
}

func (e *Payment) SetPaidAt(t time.Time) {
	e.PaidAt = &t
}

func (e *Payment) SetPaymentLink(url string) {
	e.PaymentLink = &url
}

func (e *Payment) SetPaymentDetail(paymentDetail *PaymentDetail) {
	e.PaymentDetail = paymentDetail
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
