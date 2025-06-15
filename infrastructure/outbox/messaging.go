package outbox

import (
	"context"
	"errors"

	"github.com/arvinpaundra/cent/payment/domain/outbox/constant"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type Messaging struct {
	conn *nats.Conn
}

func NewMessaging(conn *nats.Conn) Messaging {
	return Messaging{conn: conn}
}

func (r Messaging) Publish(ctx context.Context, topic string, payload []byte) error {
	js, err := jetstream.New(r.conn)
	if err != nil {
		return err
	}

	_, err = js.Stream(ctx, constant.StreamDonation)
	if err != nil && !errors.Is(err, jetstream.ErrStreamNotFound) {
		return err
	}

	if errors.Is(err, jetstream.ErrStreamNotFound) {
		_, err = js.CreateStream(ctx, jetstream.StreamConfig{
			Name:     constant.StreamDonation,
			Subjects: []string{constant.EventDonationPaid},
		})

		if err != nil {
			return err
		}
	}

	_, err = js.Publish(ctx, topic, payload)
	if err != nil {
		return err
	}

	return nil
}
