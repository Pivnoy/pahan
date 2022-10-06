package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
)

type ordersRoutes struct {
	t usecase.Order
}

func newOrdersRoutes(handler *gin.RouterGroup, t usecase.Order) {
	r := &ordersRoutes{t: t}

	handler.POST("/create_order", r.createNewOrder)
	handler.GET("/get_orders", r.getOrders)
}

type createOrderRequest struct {
	ModelID   int64  `json:"model_id"`
	Quantity  int64  `json:"quantity"`
	OrderType string `json:"order_type"`
}

func (d *createOrderRequest) validate() bool {
	return d.ModelID > 0 &&
		d.Quantity > 0 &&
		len(d.OrderType) > 0 &&
		len(d.OrderType) <= 50
}

// CreateOrder godoc
// @Summary create new order
// @Tags Posts
// @Description Create and link new order with
// @Param 		request body createOrderRequest true "query params"
// @Success     200 {object} nil
// @Failure     400 {object} errResponse
// @Router      /v1/create_order [post]
func (o *ordersRoutes) createNewOrder(c *gin.Context) {
	var request createOrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	isValid := request.validate()
	if !isValid {
		errorResponse(c, http.StatusBadRequest, "request is not valid")
		return
	}
	err := o.t.CreateOrder(c.Request.Context(),
		entity.Order{
			ModelID:   request.ModelID,
			Quantity:  request.Quantity,
			OrderType: request.OrderType,
		})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

type orderResponse struct {
	Orders []entity.Order `json:"orders"`
}

// GetOrders godoc
// @Summary get orders
// @Tags Gets
// @Description Get All orders info
// @Success     200 {array} entity.Order
// @Failure     500 {object} errResponse
// @Router      /v1/get_orders [get]
func (o *ordersRoutes) getOrders(c *gin.Context) {
	orders, err := o.t.GetOrders(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, orderResponse{orders})
}
