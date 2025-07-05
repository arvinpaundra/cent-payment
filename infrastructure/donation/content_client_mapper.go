package donation

import (
	"context"

	"github.com/arvinpaundra/cent/payment/domain/donation/external"
	contentpb"github.com/arvinpaundra/centpb/gen/go/content/v1"
)

type ContentClientMapper struct {
	client contentpb.ContentServiceClient
}

func NewContentClientMapper(client contentpb.ContentServiceClient) ContentClientMapper {
	return ContentClientMapper{
		client: client,
	}
}

func (r ContentClientMapper) FindActiveContent(ctx context.Context, userId int64) (*external.ActiveContentResponse, error) {
	req := &contentpb.FindActiveContentRequest{
		UserId: userId,
	}

	resp, err := r.client.FindActiveContent(ctx, req)
	if err != nil {
		return nil, err
	}

	result := external.ActiveContentResponse{
		ID:         resp.GetId(),
		CampaignId: resp.GetCampaignId(),
	}

	return &result, nil
}
