package httpsvc

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CustomerHTTPService :nodoc:
type CustomerHTTPService struct {
}

// NewCustomerHTTPService :nodoc:
func NewCustomerHTTPService() *CustomerHTTPService {
	return &CustomerHTTPService{}
}

func (h *CustomerHTTPService) Routes(r *gin.Engine) {
	g := r.Group("/customer")

	g.GET("", func(c *gin.Context) {
		// TODO: implement this
		c.String(http.StatusOK, "customer")
	})
}
