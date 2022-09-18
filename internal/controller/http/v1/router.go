package v1

import (
	"github.com/gin-gonic/gin"
	"pahan/internal/usecase"
)

func NewRouter(handler *gin.Engine,
	en usecase.Engine,
	su usecase.Suspension,
	ds usecase.Model,
	or usecase.Orders,
	sh usecase.Shipment,
	vd usecase.Vendor) {

	h := handler.Group("/v1")

	{
		newEngineRoutes(h, en)
		newSuspensionRoutes(h, su)
		newModelRoutes(h, ds)
		newOrdersRoutes(h, or)
		newShipmentRoutes(h, sh)
		newVendorRoutes(h, vd)
	}
}
