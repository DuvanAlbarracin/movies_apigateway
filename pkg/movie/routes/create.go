package routes

import (
	"net/http"

	"github.com/DuvanAlbarracin/movies_apigateway/pkg/movie/proto"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type CreateRequestBody struct {
	Title       *string `json:"title"`
	DirectorId  *int64  `json:"director_id"`
	ReleaseYear *int32  `json:"release_year"`
	MusicBy     *int64  `json:"music_by"`
	WrittenBy   *int64  `json:"written_by"`
}

func Create(ctx *gin.Context, c proto.MoviesServiceClient) {
	body := CreateRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if body.Title == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot create a movie without Title",
		})
		return
	}
	if body.ReleaseYear == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot create a movie without Release Year",
		})
		return
	}

	res, err := c.Create(ctx, &proto.CreateRequest{
		Title:       *body.Title,
		ReleaseYear: *body.ReleaseYear,
		DirectorID:  body.DirectorId,
		MusicBy:     body.DirectorId,
		WrittenBy:   body.WrittenBy,
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
