package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"strconv"
)

type componentRoutes struct {
	t usecase.Component
}

func newComponentRoutes(handler *gin.RouterGroup, t usecase.Component) {
	r := &componentRoutes{t: t}

	handler.GET("/get_components", r.getComponentsByVendorAndType)
}

type componentResponse struct {
	Cmp []entity.Component `json:"components"`
}

type componetsRequest struct {
	VendorID int64 `json:"vendor-id"`
	TypeID   int64 `json:"type-id"`
}

func (f *componentRoutes) getComponentsByVendorAndType(c *gin.Context) {
	vendorIDParam := c.Query("vendor-id")
	typeIDParam := c.Query("type-id")
	vendorID, err := strconv.ParseInt(vendorIDParam, 10, 64)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	typeID, err := strconv.ParseInt(typeIDParam, 10, 64)
	components, err := f.t.GetComponentsByVendorAndType(c.Request.Context(), vendorID, typeID)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, componentResponse{components})
}
