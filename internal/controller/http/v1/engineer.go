package v1

import (
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type engineerRoute struct {
	en usecase.Engineer
}

func newEngineerRoutes(handler *gin.RouterGroup, en usecase.Engineer) {
	r := engineerRoute{en: en}

	handler.GET("/engineer-by-vendor-id", r.getEngineerByIdVendor)
}

type engineerResponse struct {
	Engineer []entity.Engineer `json:"engineer"`
}

func (en *engineerRoute) getEngineerByIdVendor(c *gin.Context) {

	vendorIDString := c.Query("vendor-id")
	vendorID, _ := strconv.Atoi(vendorIDString)
	listEngineer, err := en.en.GetAllEngineerByIdVendor(c.Request.Context(), int64(vendorID))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, engineerResponse{listEngineer})
}
