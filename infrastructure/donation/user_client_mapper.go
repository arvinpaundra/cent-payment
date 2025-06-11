package donation

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/data"
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

func (r UserClientMapper) FindUserDetail(ctx context.Context, slug string) (*data.UserResponse, error) {
	req := &userpb.FindUserDetailRequest{
		Slug: slug,
	}

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("api-key", r.apiKey))

	resp, err := r.client.FindUserDetail(ctx, req)
	if err != nil {
		return nil, err
	}

	var image *string

	if resp.GetImage() != nil {
		imageStr := resp.GetImage().GetValue()
		image = &imageStr
	}

	result := data.UserResponse{
		ID:       resp.GetId(),
		Fullname: resp.GetFullname(),
		Email:    resp.GetEmail(),
		Image:    image,
	}

	return &result, nil
}
