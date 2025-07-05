package donation

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/external"
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

func (r Midtrans) Pay(ctx context.Context, pg *external.PaymentGatewayRequest) (*external.PaymentGatewayResponse, error) {
	// TODO: future we will migrate to coreapi to improvise UI/UX
	client := snap.Client{}

	client.New(r.serverKey, r.getMode())

	payload := r.parse(pg)

	res, err := client.CreateTransaction(&payload)
	if err != nil {
		return nil, err
	}

	result := external.PaymentGatewayResponse{
		Token: res.Token,
		Url:   res.RedirectURL,
	}

	return &result, nil
}

func (r Midtrans) parse(pg *external.PaymentGatewayRequest) snap.Request {
	payload := snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  pg.Code,
			GrossAmt: int64(pg.Amount),
		},
		EnabledPayments: []snap.SnapPaymentType{
			snap.PaymentTypeGopay,
			snap.PaymentTypeShopeepay,
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
