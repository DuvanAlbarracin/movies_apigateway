package routes

import (
	"net/http"

	"github.com/DuvanAlbarracin/movies_apigateway/pkg/profile/proto"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type CreateRequestBody struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Age       *int32  `json:"age"`
	Gender    *int32  `json:"gender"`
	City      *string `json:"city"`
	Role      *int32  `json:"role"`
}

func Create(ctx *gin.Context, c proto.ProfilesServiceClient) {

	body := CreateRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if body.FirstName == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot create a profile without First Name",
		})
		return
	}
	if body.Role == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot create a profile without Role",
		})
		return
	}

	res, err := c.Create(ctx, &proto.CreateRequest{
		FirstName: *body.FirstName,
		LastName:  body.LastName,
		Age:       body.Age,
		Gender:    (*proto.Gender)(body.Gender),
		City:      body.City,
		Role:      proto.Role(*body.Role),
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
