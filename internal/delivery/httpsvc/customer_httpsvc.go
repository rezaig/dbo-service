package httpsvc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezaig/dbo-service/internal/model"
)

// CustomerHTTPService :nodoc:
type CustomerHTTPService struct {
	customerUsecase model.CustomerUsecase
}

// NewCustomerHTTPService :nodoc:
func NewCustomerHTTPService(customerUsecase model.CustomerUsecase) *CustomerHTTPService {
	return &CustomerHTTPService{customerUsecase: customerUsecase}
}

func (h *CustomerHTTPService) Routes(r *gin.Engine) {
	g := r.Group("/customer")

	g.GET("", h.FindAll)
}

func (h *CustomerHTTPService) FindAll(c *gin.Context) {
	results, err := h.customerUsecase.FindAll(c.Request.Context())
	if err != nil {
		parseError(c, err)
		return
	}

	c.JSON(http.StatusOK, results)
}
