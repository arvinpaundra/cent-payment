package request

type CreateDonation struct {
	UserId int64   `json:"user_id"`
	Name   string  `json:"name" validate:"required,min=3,max=50"`
	Amount float64 `json:"amount" validate:"required"`
	Phone  string  `json:"phone"`
	Email  string  `json:"email"`
}
