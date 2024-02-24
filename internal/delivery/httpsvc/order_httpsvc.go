package httpsvc

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// OrderHTTPService :nodoc:
type OrderHTTPService struct {
}

// NewOrderHTTPService :nodoc:
func NewOrderHTTPService() *OrderHTTPService {
	return &OrderHTTPService{}
}

func (h *OrderHTTPService) Routes(r *gin.Engine) {
	g := r.Group("/order")

	g.GET("", func(c *gin.Context) {
		// TODO: implement this
		c.String(http.StatusOK, "order")
	})
}
