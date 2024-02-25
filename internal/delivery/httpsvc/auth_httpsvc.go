package httpsvc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezaig/dbo-service/internal/model"
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

	g.POST("/login", h.Login)
	g.POST("/register", h.Register)
}

func (h *AuthHTTPService) Login(c *gin.Context) {
	var bodyReq model.LoginRequest
	if err := c.Bind(&bodyReq); err != nil {
		parseError(c, err)
		return
	}

	token, err := h.authUsecase.Login(c.Request.Context(), bodyReq)
	if err != nil {
		parseError(c, err)
		return
	}

	c.JSON(http.StatusOK, model.LoginResponse{AccessToken: token})
}

func (h *AuthHTTPService) Register(c *gin.Context) {
	var bodyReq model.RegisterRequest
	if err := c.Bind(&bodyReq); err != nil {
		parseError(c, err)
		return
	}

	token, err := h.authUsecase.Register(c.Request.Context(), bodyReq)
	if err != nil {
		parseError(c, err)
		return
	}

	c.JSON(http.StatusCreated, model.LoginResponse{AccessToken: token})
}
