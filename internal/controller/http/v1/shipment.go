package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"time"
)

type shipmentRoutes struct {
	t usecase.Shipment
}

func newShipmentRoutes(handler *gin.RouterGroup, t usecase.Shipment) {
	r := &shipmentRoutes{t: t}

	handler.POST("/create_shipment", r.doNewShipment)
}

type createShipmentRequest struct {
	OrderID   int64  `json:"order_id"`
	CountryID int64  `json:"country_id"`
	Date      string `json:"date"`
}

// CreateShipment godoc
// @Summary create new shipment
// @Tags Posts
// @Description Create and link new shipment
// @Param 		request body createShipmentRequest true "query params"
// @Success     200 {object} nil
// @Failure     400 {object} errResponse
// @Failure     500 {object} errResponse
// @Router      /v1/create_shipment [post]
func (s *shipmentRoutes) doNewShipment(c *gin.Context) {
	var request createShipmentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	t, err := time.Parse(time.RFC3339, request.Date)

	if err != nil {
		errorResponse(c, http.StatusBadRequest, "error in parsing time")
		return
	}

	err = s.t.CreateShipment(c.Request.Context(),
		entity.Shipment{
			OrderID:     request.OrderID,
			CountryToID: request.CountryID,
			Date:        t,
		})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
