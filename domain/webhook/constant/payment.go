package constant

type PaymentSource string

func (c PaymentSource) String() string {
	return string(c)
}

const (
	PaymentSourceMidtrans PaymentSource = "midtrans"
	PaymentSourceOthers   PaymentSource = "others"
)

type PaymentStatus string

func (c PaymentStatus) String() string {
	return string(c)
}

const (
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusPaid    PaymentStatus = "paid"
	PaymentStatusFailed  PaymentStatus = "failed"
	PaymentStatusExpired PaymentStatus = "expired"
)

type PaymentPurpose string

func (c PaymentPurpose) String() string {
	return string(c)
}

const (
	PaymentPurposeDonation PaymentPurpose = "donation"
	PaymentPurposeOthers   PaymentPurpose = "others"
)

type PaymentMethod string

func (c PaymentMethod) String() string {
	return string(c)
}

const (
	PaymentMethodGopay     PaymentMethod = "gopay"
	PaymentMethodShopeepay PaymentMethod = "shopeepay"
	PaymentMethodQris      PaymentMethod = "qris"
	PaymentMethodNone      PaymentMethod = "none"
	PaymentMethodOthers    PaymentMethod = "others"
)
