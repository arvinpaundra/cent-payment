package donation

type CreateDonation struct {
	UserSlug string  `json:"-"`
	Name     string  `json:"name" validate:"required,min=3,max=50"`
	Amount   float64 `json:"amount" validate:"required"`
	Message  string  `json:"message" validate:"max=250"`
	Phone    *string `json:"phone"`
	Email    *string `json:"email"`
}
