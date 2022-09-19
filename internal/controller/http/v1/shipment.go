package v1

import (
	"github.com/gin-gonic/gin"
	"log"
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

	handler.POST("/accept_order", r.doNewShipment)
}

type doShipmentRequest struct {
	OrderID   int64  `json:"order_id"`
	CountryID int64  `json:"country_id"`
	Date      string `json:"date"`
}

func (s *shipmentRoutes) doNewShipment(c *gin.Context) {
	var request doShipmentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Fatal(err)
		return
	}

	t, _ := time.Parse(time.RFC3339, request.Date)

	err := s.t.NewShipment(c.Request.Context(),
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
