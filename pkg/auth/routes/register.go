package routes

import (
	"context"
	"errors"
	"net/http"

	"github.com/DuvanAlbarracin/movies_apigateway/pkg/auth/proto"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type RegisterRequestBody struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirmation"`
}

func Register(ctx *gin.Context, c proto.AuthServiceClient) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if !body.passwordConfirmationCheck() {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Passwords must match"))
		return
	}

	res, err := c.Register(context.Background(), &proto.RegisterRequest{
		Username: utils.SanitizeString(body.Username),
		Email:    utils.SanitizeString(body.Email),
		Password: body.Password,
	})

	if err != nil {
		rpcCode := utils.RpcCode{Code: status.Convert(err).Code()}
		ctx.JSON(int(rpcCode.ToHttpCode()), gin.H{
			"error": status.Convert(err).Message(),
		})
		return
	}

	ctx.JSON(int(res.Status), &res)
}

func (body RegisterRequestBody) passwordConfirmationCheck() bool {
	return body.Password == body.PasswordConfirm
}
