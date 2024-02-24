package httpsvc

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthHTTPService :nodoc:
type AuthHTTPService struct {
}

// NewAuthHTTPService :nodoc:
func NewAuthHTTPService() *AuthHTTPService {
	return &AuthHTTPService{}
}

func (h *AuthHTTPService) Routes(r *gin.Engine) {
	g := r.Group("/auth")

	g.GET("/login", func(c *gin.Context) {
		// TODO: implement this
		c.String(http.StatusOK, "login")
	})
}
