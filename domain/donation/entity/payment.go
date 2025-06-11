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
	Type        constant.PaymentType
	Status      constant.PaymentStatus
	Method      constant.PaymentMethod
	Amount      float64
	Currency    *string
	BankName    *string
	VaNumber    *string
	Qrcode      *string
	PaymentLink *string
	ExpiredAt   *time.Time

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

func (e *Payment) IsBankTransfer() bool {
	return e.Method == constant.PaymentMethodBankTransfer
}

func (e *Payment) SetBankName(bankName string) {
	e.BankName = &bankName
}

func (e *Payment) SetPaymentLink(url string) {
	e.PaymentLink = &url
}

func (e *Payment) SetPaymentDetail(pd *PaymentDetail) {
	e.PaymentDetail = &PaymentDetail{
		ContentId:  pd.ContentId,
		Name:       pd.Name,
		Message:    pd.Message,
		Phone:      pd.Phone,
		Email:      pd.Email,
		CampaignId: pd.CampaignId,
	}
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
