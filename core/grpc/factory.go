package grpc

import (
	"log"

	"github.com/arvinpaundra/centpb/gen/go/auth/v1"
	"github.com/arvinpaundra/centpb/gen/go/content/v1"
	"github.com/arvinpaundra/centpb/gen/go/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type usersvc struct {
	auth.AuthenticateServiceClient
	user.UserServiceClient
}

type contentsvc struct {
	content.ContentServiceClient
}

type ClientConfig struct {
	UserClientAddr    string
	ContentClientAddr string
}

type Client struct {
	userClient    usersvc
	contentClient contentsvc

	conns []*grpc.ClientConn
}

func NewClient(config ClientConfig) *Client {
	client := new(Client)

	userConn, err := client.dial(config.UserClientAddr)
	if err != nil {
		log.Fatalf("failed to dial to client: %v\n", err.Error())
	}

	client.userClient = usersvc{
		AuthenticateServiceClient: auth.NewAuthenticateServiceClient(userConn),
		UserServiceClient:         user.NewUserServiceClient(userConn),
	}

	contentConn, err := client.dial(config.ContentClientAddr)
	if err != nil {
		log.Fatalf("failed to dial to client: %v\n", err.Error())
	}

	client.contentClient = contentsvc{
		ContentServiceClient: content.NewContentServiceClient(contentConn),
	}

	return client
}

func (c *Client) Close() error {
	for _, conn := range c.conns {
		err := conn.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) UserClient() usersvc {
	return c.userClient
}

func (c *Client) ContentClient() contentsvc {
	return c.contentClient
}

func (c *Client) appendConn(conn *grpc.ClientConn) {
	c.conns = append(c.conns, conn)
}

func (c *Client) dial(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c.appendConn(conn)

	return conn, nil
}
