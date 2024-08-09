package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/DuvanAlbarracin/movies_apigateway/pkg/genre/proto"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type RemoveFromMovieRequestBody struct {
	GenreID int64 `json:"genre_id"`
	MovieID int64 `json:"movie_id"`
}

func RemoveFromMovie(ctx *gin.Context, c proto.GenreServiceClient) {
	genre_id, err := strconv.Atoi(ctx.Param("genre_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot parse genre_id param",
		})
		return
	}

	movie_id, err := strconv.Atoi(ctx.Param("movie_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot parse movie_id param",
		})
		return
	}

	res, err := c.RemoveFromMovie(context.Background(), &proto.RemoveFromMovieRequest{
		GenreId: int64(genre_id),
		MovieId: int64(movie_id),
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
