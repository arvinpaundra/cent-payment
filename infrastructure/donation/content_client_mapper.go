package donation

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/data"
	"github.com/arvinpaundra/centpb/gen/go/content/v1"
)

type ContentClientMapper struct {
	client content.ContentServiceClient
}

func NewContentClientMapper(client content.ContentServiceClient) ContentClientMapper {
	return ContentClientMapper{
		client: client,
	}
}

func (r ContentClientMapper) FindActiveContent(ctx context.Context, userId int64) (*data.ActiveContentResponse, error) {
	req := &content.FindActiveContentRequest{
		UserId: userId,
	}

	resp, err := r.client.FindActiveContent(ctx, req)
	if err != nil {
		return nil, err
	}

	result := data.ActiveContentResponse{
		ID:         resp.GetId(),
		CampaignId: resp.GetCampaignId(),
	}

	return &result, nil
}
