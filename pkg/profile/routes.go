package profile

import (
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/config"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/profile/routes"
	"github.com/gin-gonic/gin"
)

// func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	// aut := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routeGroup := r.Group("/profiles")
	// routes.Use(aut.AuthRequired)

	routeGroup.POST("/create", svc.Create)
	routeGroup.DELETE("/delete/:id", svc.Delete)
	routeGroup.PATCH("/modify/:id", svc.Modify)
	routeGroup.GET("/:id", svc.GetById)
	routeGroup.GET("", svc.GetAll)

	return svc
}

func (svc *ServiceClient) Create(ctx *gin.Context) {
	routes.Create(ctx, svc.Client)
}

func (svc *ServiceClient) Delete(ctx *gin.Context) {
	routes.Delete(ctx, svc.Client)
}

func (svc *ServiceClient) Modify(ctx *gin.Context) {
	routes.Modify(ctx, svc.Client)
}

func (svc *ServiceClient) GetById(ctx *gin.Context) {
	routes.GetById(ctx, svc.Client)
}

func (svc *ServiceClient) GetAll(ctx *gin.Context) {
	routes.GetAll(ctx, svc.Client)
}
