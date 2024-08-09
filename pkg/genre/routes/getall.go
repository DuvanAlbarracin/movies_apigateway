package routes

import (
	"context"
	"net/http"

	"github.com/DuvanAlbarracin/movies_apigateway/pkg/genre/proto"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

func GetAll(ctx *gin.Context, c proto.GenreServiceClient) {
	res, err := c.GetAll(context.Background(), &proto.GetAllRequest{})

	if err != nil {
		rpcCode := utils.RpcCode{Code: status.Convert(err).Code()}
		ctx.JSON(int(rpcCode.ToHttpCode()), gin.H{
			"error": status.Convert(err).Message(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
