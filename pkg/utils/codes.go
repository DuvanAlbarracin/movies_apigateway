package utils

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type RpcCode struct {
	Code codes.Code
}

func (c *RpcCode) ToHttpCode() (httpCode int64) {
	switch int(c.Code) {
	case 0:
		httpCode = http.StatusOK

	case 3, 6, 9:
		httpCode = http.StatusBadRequest

	case 4:
		httpCode = http.StatusRequestTimeout

	case 5:
		httpCode = http.StatusNotFound

	case 7, 16:
		httpCode = http.StatusUnauthorized

	case 12:
		httpCode = http.StatusNotImplemented

	case 13:
		httpCode = http.StatusInternalServerError

	case 14:
		httpCode = http.StatusServiceUnavailable

	default:
		httpCode = http.StatusInternalServerError
	}

	return
}
