package external

type PaymentGatewayRequest struct {
	Amount float64
	Code   string
}

type PaymentGatewayResponse struct {
	Token string
	Url   string
}
