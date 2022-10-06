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

	handler.GET("/get_components", r.getComponents)
	handler.GET("/get_components_by_vendor_and_type", r.getComponentsByVendorAndType)
}

type componentResponse struct {
	Cmp []entity.Component `json:"components"`
}

// GetComponentByVendorIDAndTypeID godoc
// @Summary get component
// @Tags Gets
// @Description Get all components depend on vendorID and typeID
// @Param       vendor-id  query   string  true "id of a vendor"
// @Param       type-id    query   string  true "id of a type"
// @Success     200 {array}  entity.Component
// @Failure     400 {object} errResponse
// @Router      /v1/get_components_by_vendor_and_type [get]
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

type componentsResponse struct {
	Components []entity.Component `json:"components"`
}

// GetComponents godoc
// @Summary get components
// @Tags Gets
// @Description Get all components
// @Success     200 {array}  entity.Component
// @Failure     500 {object} errResponse
// @Router      /v1/get_components [get]
func (f *componentRoutes) getComponents(c *gin.Context) {
	components, err := f.t.GetComponents(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, componentsResponse{components})
}
