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
	var params model.CustomerParams
	if err := c.BindQuery(&params); err != nil {
		parseError(c, err)
		return
	}
	if err := params.Validate(); err != nil {
		parseError(c, err)
		return
	}

	results, totalItems, err := h.customerUsecase.FindAll(c.Request.Context(), params)
	if err != nil {
		parseError(c, err)
		return
	}

	paginationResponse := model.PaginationResponse{
		Pagination: model.Pagination{
			Page:       params.Page,
			PerPage:    params.PerPage,
			TotalItems: totalItems,
		},
		Data: results,
	}

	c.JSON(http.StatusOK, paginationResponse)
}
