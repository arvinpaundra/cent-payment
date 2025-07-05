package handler

import (
	"github.com/arvinpaundra/cent/payment/core/grpc"
	"github.com/arvinpaundra/cent/payment/core/validator"
	"gorm.io/gorm"
)

type Handler struct {
	db         *gorm.DB
	grpcClient *grpc.Client
	validator  *validator.Validator
}

func NewHandler(db *gorm.DB, grpcClient *grpc.Client, validator *validator.Validator) Handler {
	return Handler{
		db:         db,
		grpcClient: grpcClient,
		validator:  validator,
	}
}
