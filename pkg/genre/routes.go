package genre

import (
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/config"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/genre/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routeGroup := r.Group("/genres")

	routeGroup.GET("/:id", svc.GetById)
	routeGroup.GET("", svc.GetAll)
	routeGroup.POST("/:id", svc.AddToMovie)
	routeGroup.DELETE("", svc.RemoveFromMovie)

	return svc
}

func (svc *ServiceClient) GetById(ctx *gin.Context) {
	routes.GetById(ctx, svc.Client)
}

func (svc *ServiceClient) GetAll(ctx *gin.Context) {
	routes.GetAll(ctx, svc.Client)
}

func (svc *ServiceClient) AddToMovie(ctx *gin.Context) {
	routes.AddToMovie(ctx, svc.Client)
}

func (svc *ServiceClient) RemoveFromMovie(ctx *gin.Context) {
	routes.RemoveFromMovie(ctx, svc.Client)
}
