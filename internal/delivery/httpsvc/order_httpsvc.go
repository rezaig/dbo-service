package httpsvc

import (
	"github.com/gin-gonic/gin"
	"github.com/rezaig/dbo-service/internal/model"
	"net/http"
)

// OrderHTTPService :nodoc:
type OrderHTTPService struct {
	orderUsecase model.OrderUsecase
}

// NewOrderHTTPService :nodoc:
func NewOrderHTTPService(orderUsecase model.OrderUsecase) *OrderHTTPService {
	return &OrderHTTPService{orderUsecase: orderUsecase}
}

func (h *OrderHTTPService) Routes(r *gin.Engine) {
	g := r.Group("/order")

	g.GET("", func(c *gin.Context) {
		// TODO: implement this
		c.String(http.StatusOK, "order")
	})
}
