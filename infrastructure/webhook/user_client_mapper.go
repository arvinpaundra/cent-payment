package webhook

import (
	"context"
	"errors"

	"github.com/arvinpaundra/cent/payment/domain/webhook/data"
	userpb "github.com/arvinpaundra/centpb/gen/go/user/v1"
	"google.golang.org/grpc/metadata"
)

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

func (r UserClientMapper) UpdateUserBalance(ctx context.Context, payload *data.UpdateBalanceUserRequest) error {
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
