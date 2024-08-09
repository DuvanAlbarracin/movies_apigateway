package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/DuvanAlbarracin/movies_apigateway/pkg/movie/proto"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type GetByIdRequestBody struct {
	Id int64 `json:"id"`
}

func GetById(ctx *gin.Context, c proto.MoviesServiceClient) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot parse Id param",
		})
		return
	}

	res, err := c.GetById(context.Background(), &proto.GetByIdRequest{
		Id: int64(id),
	})
	if err != nil {
		rpcCode := utils.RpcCode{Code: status.Convert(err).Code()}
		ctx.JSON(int(rpcCode.ToHttpCode()), gin.H{
			"error": status.Convert(err).Message(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
