package routes

import (
	"net/http"

	"github.com/DuvanAlbarracin/movies_apigateway/pkg/profile/proto"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type ModifyRequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int32  `json:"age"`
	Gender    int32  `json:"gender"`
	City      string `json:"city"`
	Role      int32  `json:"role"`
}

func Modify(ctx *gin.Context, c proto.ProfilesServiceClient) {
	body := ModifyRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	body.FirstName = utils.SanitizeString(body.FirstName)
	body.LastName = utils.SanitizeString(body.LastName)
	body.City = utils.SanitizeString(body.City)

	res, err := c.Modify(ctx, &proto.ModifyRequest{
		FirstName: &body.FirstName,
		LastName:  &body.LastName,
		Age:       &body.Age,
		Gender:    (*proto.Gender)(&body.Gender),
		City:      &body.City,
		Role:      (*proto.Role)(&body.Role),
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
