package donation

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/entity"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

const (
	ModeSandbox    = "sandbox"
	ModeProduction = "production"
)

type Midtrans struct {
	merchantId string
	clientKey  string
	serverKey  string

	mode string
}

func NewMidtrans(serverKey, mode string) Midtrans {
	return Midtrans{
		serverKey: serverKey,
		mode:      mode,
	}
}

func (r Midtrans) Pay(ctx context.Context, pg *entity.PaymentGateway) (string, error) {
	// TODO: future we will migrate to coreapi to improvise UI/UX
	client := snap.Client{}

	client.New(r.serverKey, r.getMode())

	payload := r.parse(pg)

	url, err := client.CreateTransactionUrl(&payload)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (r Midtrans) parse(pg *entity.PaymentGateway) snap.Request {
	payload := snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  pg.Code,
			GrossAmt: int64(pg.Amount),
		},
		EnabledPayments: []snap.SnapPaymentType{
			snap.PaymentTypeGopay,
			snap.PaymentTypeShopeepay,
			snap.PaymentTypeBCAVA,
			snap.PaymentTypeBRIVA,
			snap.PaymentTypeBNIVA,
			snap.PaymentTypePermataVA,
			snap.PaymentTypeEChannel,
		},
		Expiry: &snap.ExpiryDetails{
			Unit:     "minute",
			Duration: 15,
		},
	}

	return payload
}

func (r Midtrans) getMode() midtrans.EnvironmentType {
	if r.mode == ModeProduction {
		return midtrans.Production
	}

	return midtrans.Sandbox
}
