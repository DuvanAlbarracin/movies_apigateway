package profile

import (
	"log"

	"github.com/DuvanAlbarracin/movies_api_gateway/pkg/config"
	"github.com/DuvanAlbarracin/movies_api_gateway/pkg/profile/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client proto.ProfilesServiceClient
}

func InitServiceClient(c *config.Config) proto.ProfilesServiceClient {
	cc, err := grpc.Dial(c.ProfileUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Could not connect with ProfileServiceClient:", err)
	}

	return proto.NewProfilesServiceClient(cc)
}
