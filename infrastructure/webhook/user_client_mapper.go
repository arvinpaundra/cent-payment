package webhook

import (
	"context"
	"errors"

	"github.com/arvinpaundra/cent/payment/domain/webhook/external"
	"github.com/arvinpaundra/cent/payment/domain/webhook/repository"
	userpb "github.com/arvinpaundra/centpb/gen/go/user/v1"
	"google.golang.org/grpc/metadata"
)

var _ repository.UserClientMapper = (*UserClientMapper)(nil)

type UserClientMapper struct {
	client userpb.UserServiceClient
	apiKey string
}

func NewUserClientMapper(
	client userpb.UserServiceClient,
	apiKey string,
) UserClientMapper {
	return UserClientMapper{
		client: client,
		apiKey: apiKey,
	}
}

func (r UserClientMapper) FindUserDetail(ctx context.Context, userId int64) (*external.UserResponse, error) {
	req := &userpb.FindUserDetailRequest{
		Id: userId,
	}

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("api-key", r.apiKey))

	resp, err := r.client.FindUserDetail(ctx, req)
	if err != nil {
		return nil, err
	}

	result := external.UserResponse{
		ID:       resp.User.GetId(),
		Fullname: resp.User.GetFullname(),
		Email:    resp.User.GetEmail(),
		Key:      resp.User.GetKey(),
		Image:    resp.User.Image,
	}

	return &result, nil
}

func (r UserClientMapper) UpdateUserBalance(ctx context.Context, payload *external.UpdateBalanceUserRequest) error {
	req := userpb.UpdateUserBalanceRequest{
		Id:     payload.UserId,
		Amount: payload.Amount,
	}

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("api-key", r.apiKey))

	res, err := r.client.UpdateUserBalance(ctx, &req)
	if err != nil {
		return errors.New(res.Meta.GetMessage())
	}

	return nil
}
