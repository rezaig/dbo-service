package httpsvc

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezaig/dbo-service/internal/model"
)

type OrderHTTPService struct {
	orderUsecase model.OrderUsecase
}

func NewOrderHTTPService(orderUsecase model.OrderUsecase) *OrderHTTPService {
	return &OrderHTTPService{orderUsecase: orderUsecase}
}

func (h *OrderHTTPService) Routes(r *gin.Engine) {
	g := r.Group("/order")

	g.Use(AuthMiddleware())

	g.GET("", h.FindAll)
}

func (h *OrderHTTPService) FindAll(c *gin.Context) {
	claims, exists := c.Get(model.ClaimsCtxKey)
	if !exists {
		parseError(c, errors.New("claim does not exist"))
		return
	}

	claim, ok := claims.(*model.CustomClaims)
	if !ok {
		parseError(c, errors.New("invalid claim"))
		return
	}

	var params model.OrderParams
	if err := c.BindQuery(&params); err != nil {
		parseError(c, err)
		return
	}
	if err := params.Validate(); err != nil {
		parseError(c, err)
		return
	}

	results, totalItems, err := h.orderUsecase.FindAllByCustomerID(c.Request.Context(), claim.AccountID, params)
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
