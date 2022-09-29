package v1

import (
	"github.com/gin-gonic/gin"
	"pahan/internal/usecase"
)

func NewRouter(handler *gin.Engine,
	ds usecase.Model,
	or usecase.Order,
	sh usecase.Shipment,
	vd usecase.Vendor) {

	h := handler.Group("/v1")

	{
		newModelRoutes(h, ds)
		newOrdersRoutes(h, or)
		newShipmentRoutes(h, sh)
		newVendorRoutes(h, vd)
	}
}
