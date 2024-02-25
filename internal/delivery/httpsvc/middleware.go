package httpsvc

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rezaig/dbo-service/internal/helper"
	"github.com/rezaig/dbo-service/internal/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, errorResponse{Error: "unauthorized"})
			c.Abort()
			return
		}

		authSplit := strings.Fields(authHeader)
		if len(authSplit) != 2 || authSplit[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, errorResponse{Error: "unauthorized"})
			c.Abort()
			return
		}

		jwtToken := authSplit[1]
		claims, err := helper.DecodeJWTToken(jwtToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, errorResponse{Error: "unauthorized"})
			c.Abort()
			return
		}

		c.Set(model.ClaimsCtxKey, claims)
		c.Next()
	}
}
