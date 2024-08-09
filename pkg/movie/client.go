package movie

import (
	"log"

	"github.com/DuvanAlbarracin/movies_apigateway/pkg/config"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/movie/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client proto.MoviesServiceClient
}

func InitServiceClient(c *config.Config) proto.MoviesServiceClient {
	cc, err := grpc.Dial(c.MovieUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Could not connect with MovieServiceClient:", err)
	}

	return proto.NewMoviesServiceClient(cc)
}
