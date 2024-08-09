package routes

import (
	"net/http"
	"strconv"

	"github.com/DuvanAlbarracin/movies_apigateway/pkg/movie/proto"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type ModifyRequestBody struct {
	Title       *string `json:"title"`
	DirectorId  *int64  `json:"director_id"`
	ReleaseYear *int32  `json:"release_year"`
	MusicBy     *int64  `json:"music_by"`
	WrittenBy   *int64  `json:"written_by"`
}

func Modify(ctx *gin.Context, c proto.MoviesServiceClient) {
	var title string

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

	if body.Title != nil {
		title = utils.SanitizeString(*body.Title)
		body.Title = &title
	}

	res, err := c.Modify(ctx, &proto.ModifyRequest{
		Id:          int64(id),
		Title:       body.Title,
		DirectorID:  body.DirectorId,
		ReleaseYear: body.ReleaseYear,
		MusicBy:     body.MusicBy,
		WrittenBy:   body.WrittenBy,
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
