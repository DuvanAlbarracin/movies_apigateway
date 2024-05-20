package auth

import (
	"fmt"

	"github.com/DuvanAlbarracin/movies_api_gateway/pkg/auth/proto"
	"github.com/DuvanAlbarracin/movies_api_gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client proto.AuthServiceClient
}

func InitServiceClient(c *config.Config) proto.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect with AuthServiceClient:", err)
	}

	return proto.NewAuthServiceClient(cc)
}
