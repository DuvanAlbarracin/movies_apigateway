package routes

import (
	"net/http"

	"github.com/DuvanAlbarracin/movies_apigateway/pkg/profile/proto"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type CreateRequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int16  `json:"age"`
	Gender    int8   `json:"gender"`
	City      string `json:"city"`
	Role      int8   `json:"role"`
}

func Create(ctx *gin.Context, c proto.ProfilesServiceClient) {
	body := CreateRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := c.Create(ctx, &proto.CreateRequest{
		FirstName: utils.SanitizeString(body.FirstName),
		LastName:  utils.SanitizeString(body.LastName),
		Age:       int32(body.Age),
		Gender:    proto.Gender(body.Gender),
		City:      utils.SanitizeString(body.City),
		Role:      proto.Role(body.Role),
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
