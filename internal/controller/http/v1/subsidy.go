package v1

import (
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"

	"github.com/gin-gonic/gin"
)

type subsidyRoutes struct {
	t usecase.Subsidy
}

func newSubsidyRoutes(handler *gin.RouterGroup, t usecase.Subsidy) {
	r := subsidyRoutes{t: t}

	handler.GET("/subsidies", r.getSubsidies)
	handler.POST("create_subsidy", r.createSubsidy)
}

type subsidyResponse struct {
	Subsidy []entity.Subsidy `json:"subsidy"`
}

func (s *subsidyRoutes) getSubsidies(c *gin.Context) {
	listSubsidies, err := s.t.GetAllSubsidies(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, subsidyResponse{listSubsidies})
}

func (s *subsidyRoutes) createSubsidy(c *gin.Context) {

}
