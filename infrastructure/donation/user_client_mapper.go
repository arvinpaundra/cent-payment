package donation

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/external"
	"github.com/arvinpaundra/cent/payment/domain/donation/repository"
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

func (r UserClientMapper) FindUserBySlug(ctx context.Context, slug string) (*external.UserResponse, error) {
	req := &userpb.FindUserBySlugRequest{
		Slug: slug,
	}

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("api-key", r.apiKey))

	resp, err := r.client.FindUserBySlug(ctx, req)
	if err != nil {
		return nil, err
	}

	result := external.UserResponse{
		ID:       resp.User.GetId(),
		Fullname: resp.User.GetFullname(),
		Email:    resp.User.GetEmail(),
		Image:    resp.User.Image,
	}

	return &result, nil
}
