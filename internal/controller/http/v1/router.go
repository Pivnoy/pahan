package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"pahan/internal/usecase"
)

func NewRouter(handler *gin.Engine,
	ds usecase.Model,
	or usecase.Order,
	sh usecase.Shipment,
	vd usecase.Vendor,
	sb usecase.Subsidy,
	en usecase.Engineer,
	fc usecase.Factory,
	cm usecase.Component,
	tp usecase.Type) {

	h := handler.Group("/v1")

	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	{
		newModelRoutes(h, ds)
		newOrdersRoutes(h, or)
		newShipmentRoutes(h, sh)
		newVendorRoutes(h, vd)
		newSubsidyRoutes(h, sb)
		newEngineerRoutes(h, en)
		newFactoryRoutes(h, fc)
		newComponentRoutes(h, cm)
		newTypeRouter(h, tp)
	}
}
