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
	PaymentStatusSuccedd PaymentStatus = "succeed"
	PaymentStatusFailed  PaymentStatus = "failed"
	PaymentStatusExpired PaymentStatus = "expired"
)

type PaymentType string

func (c PaymentType) String() string {
	return string(c)
}

const (
	PaymentTypeDonation PaymentType = "donation"
	PaymentTypeOthers   PaymentType = "others"
)

type PaymentMethod string

func (c PaymentMethod) String() string {
	return string(c)
}

const (
	PaymentMethodEmpty        PaymentMethod = ""
	PaymentMethodGopay        PaymentMethod = "gopay"
	PaymentMethodShopeepay    PaymentMethod = "shopeepay"
	PaymentMethodQris         PaymentMethod = "qris"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
	PaymentMethodNone         PaymentMethod = "none"
	PaymentMethodOthers       PaymentMethod = "others"
)

type Bank string

func (c Bank) String() string {
	return string(c)
}

const (
	BankEmpty Bank = ""
	BankBri   Bank = "bri"
	BankBca   Bank = "bca"
	BankBni   Bank = "bni"
	BankCimb  Bank = "cimb"
)
