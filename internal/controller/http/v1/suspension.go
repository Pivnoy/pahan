package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
)

type suspensionRoutes struct {
	t usecase.Suspension
}

func newSuspensionRoutes(handler *gin.RouterGroup, t usecase.Suspension) {
	r := suspensionRoutes{t: t}

	handler.GET("/suspension", r.getSuspension)
}

type suspensionResponse struct {
	Suspension []entity.Suspension `json:"suspension"`
}

func (r *suspensionRoutes) getSuspension(c *gin.Context) {
	listSuspension, err := r.t.Suspensions(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, suspensionResponse{listSuspension})
}
