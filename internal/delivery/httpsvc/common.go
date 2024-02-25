package httpsvc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Error string `json:"error"`
}

type successResponse struct {
	Success bool `json:"success"`
}

func parseError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
}
