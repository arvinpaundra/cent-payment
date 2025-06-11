package rest

import (
	"github.com/arvinpaundra/cent/payment/core/grpc"
	"github.com/arvinpaundra/cent/payment/core/validator"
	"gorm.io/gorm"
)

type Controller struct {
	db         *gorm.DB
	grpcClient *grpc.Client
	validator  *validator.Validator
}

func NewController(db *gorm.DB, grpcClient *grpc.Client, validator *validator.Validator) Controller {
	return Controller{
		db:         db,
		grpcClient: grpcClient,
		validator:  validator,
	}
}
