package httpsvc

import (
	"github.com/gin-gonic/gin"
	"github.com/rezaig/dbo-service/internal/model"
	"net/http"
)

// AuthHTTPService :nodoc:
type AuthHTTPService struct {
	authUsecase model.AuthUsecase
}

// NewAuthHTTPService :nodoc:
func NewAuthHTTPService(authUsecase model.AuthUsecase) *AuthHTTPService {
	return &AuthHTTPService{authUsecase: authUsecase}
}

func (h *AuthHTTPService) Routes(r *gin.Engine) {
	g := r.Group("/auth")

	g.GET("/login", func(c *gin.Context) {
		// TODO: implement this
		c.String(http.StatusOK, "login")
	})
}
