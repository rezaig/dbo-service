package httpsvc

import (
	"net/http"
	"strconv"

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
	g.GET("/:id", h.FindByID)
	g.POST("", h.Insert)
	g.PUT("/:id", h.Update)
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

func (h *CustomerHTTPService) FindByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		parseError(c, err)
		return
	}

	result, err := h.customerUsecase.FindByID(c.Request.Context(), int64(idInt))
	if err != nil {
		parseError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *CustomerHTTPService) Update(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		parseError(c, err)
		return
	}

	var bodyReq model.CustomerRequest
	if err := c.Bind(&bodyReq); err != nil {
		parseError(c, err)
		return
	}

	data := model.Customer{
		Name:        bodyReq.Name,
		Email:       bodyReq.Email,
		PhoneNumber: bodyReq.PhoneNumber,
	}
	updatedData, err := h.customerUsecase.Update(c.Request.Context(), data, int64(idInt))
	if err != nil {
		parseError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedData)
}

func (h *CustomerHTTPService) Insert(c *gin.Context) {
	var bodyReq model.CustomerRequest
	if err := c.Bind(&bodyReq); err != nil {
		parseError(c, err)
		return
	}

	data := model.Customer{
		Name:        bodyReq.Name,
		Email:       bodyReq.Email,
		PhoneNumber: bodyReq.PhoneNumber,
	}
	insertedData, err := h.customerUsecase.Insert(c.Request.Context(), data)
	if err != nil {
		parseError(c, err)
		return
	}

	c.JSON(http.StatusCreated, insertedData)
}

func (h *CustomerHTTPService) Delete(c *gin.Context) {
	// TODO: Implement this
}
