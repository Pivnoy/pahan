package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
)

type vendorRoutes struct {
	t usecase.Vendor
}

func newVendorRoutes(handler *gin.RouterGroup, t usecase.Vendor) {
	r := vendorRoutes{t: t}

	handler.GET("/get_all_vendors", r.getVendors)
}

type vendorResponse struct {
	Vendor []entity.Vendor `json:"vendor"`
}

func (v *vendorRoutes) getVendors(c *gin.Context) {
	listVendors, err := v.t.GetAllVendors(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, vendorResponse{listVendors})
}
