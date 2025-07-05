package external

type UserResponse struct {
	ID       int64
	Fullname string
	Email    string
	Key      string
	Image    *string
}

type UpdateBalanceUserRequest struct {
	UserId int64
	Amount float64
}
