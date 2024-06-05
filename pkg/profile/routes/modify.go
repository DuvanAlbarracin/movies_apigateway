package routes

import (
	"net/http"
	"strconv"

	"github.com/DuvanAlbarracin/movies_apigateway/pkg/profile/proto"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type ModifyRequestBody struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Age       *int32  `json:"age"`
	Gender    *int32  `json:"gender"`
	City      *string `json:"city"`
	Role      *int32  `json:"role"`
}

func Modify(ctx *gin.Context, c proto.ProfilesServiceClient) {
	var firstName string
	var lastName string
	var city string

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot parse Id param",
		})
		return
	}

	body := ModifyRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if body.FirstName != nil {
		firstName = utils.SanitizeString(*body.FirstName)
		body.FirstName = &firstName
	}

	if body.LastName != nil {
		lastName = utils.SanitizeString(*body.LastName)
		body.LastName = &lastName
	}

	if body.City != nil {
		city = utils.SanitizeString(*body.City)
		body.City = &city
	}

	res, err := c.Modify(ctx, &proto.ModifyRequest{
		Id:        int64(id),
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Age:       body.Age,
		Gender:    (*proto.Gender)(body.Gender),
		City:      body.City,
		Role:      (*proto.Role)(body.Role),
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
