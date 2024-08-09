package genre

import (
	"log"

	"github.com/DuvanAlbarracin/movies_apigateway/pkg/config"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/genre/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client proto.GenreServiceClient
}

func InitServiceClient(c *config.Config) proto.GenreServiceClient {
	cc, err := grpc.Dial(c.MovieUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Could not connect with GenreServiceClient:", err)
	}

	return proto.NewGenreServiceClient(cc)
}
