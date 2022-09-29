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

	handler.POST("/new_order", r.doNewOrder)
}

type doOrderRequest struct {
	ModelID   int64  `json:"model_id"`
	Quantity  int64  `json:"quantity"`
	OrderType string `json:"order_type"`
}

func (o *ordersRoutes) doNewOrder(c *gin.Context) {
	var request doOrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := o.t.NewOrder(c.Request.Context(),
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
