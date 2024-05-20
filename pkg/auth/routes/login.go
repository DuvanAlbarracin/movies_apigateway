package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DuvanAlbarracin/movies_apigateway/pkg/auth/proto"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c proto.AuthServiceClient) {
	body := LoginRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := c.Login(context.Background(), &proto.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		rpcCode := utils.RpcCode{Code: status.Convert(err).Code()}
		ctx.JSON(int(rpcCode.ToHttpCode()), gin.H{
			"error": status.Convert(err).Message(),
		})
		return
	}

	fmt.Println("TOKENN API:", res.GetToken())
	fmt.Println("STATUS API", res.GetStatus())

	ctx.JSON(http.StatusOK, &res)
}
