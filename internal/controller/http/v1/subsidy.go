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

	handler.GET("/get_subsidies", r.getSubsidies)
	handler.POST("/create_subsidy", r.createSubsidy)
}

type subsidyResponse struct {
	Subsidy []entity.Subsidy `json:"subsidy"`
}

// GetSubsidies godoc
// @Summary list of subsidies
// @Description Get all subsidies
// @Success     200 {array}  entity.Subsidy
// @Failure     500 {object} errResponse
// @Router      /v1/get_subsidies [get]
func (s *subsidyRoutes) getSubsidies(c *gin.Context) {
	listSubsidies, err := s.t.GetAllSubsidies(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, subsidyResponse{listSubsidies})
}

type createSubsidyRequest struct {
	CountryIDBy    int64   `json:"country-id-by"`
	RequirePriceBy float64 `json:"require-price-by"`
	RequiredWdBy   string  `json:"required-wd-by"`
}

func (c *createSubsidyRequest) validate() bool {
	return len(c.RequiredWdBy) != 3 &&
		c.CountryIDBy > 0 &&
		c.RequirePriceBy > 0
}

// CreateSubsidy godoc
// @Summary create subsidy
// @Description Create and link subsidy with dependent values
// @Param 		request body createSubsidyRequest true "query params"
// @Success     200 {object} nil
// @Failure     400 {object} errResponse
// @Router      /v1/create_subsidy [post]
func (s *subsidyRoutes) createSubsidy(c *gin.Context) {
	var sbs createSubsidyRequest
	if err := c.ShouldBindJSON(&sbs); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	isValid := sbs.validate()
	if !isValid {
		errorResponse(c, http.StatusBadRequest, "error in validating subsidy request")
		return
	}
	err := s.t.CreateSubsidy(c.Request.Context(), sbs.CountryIDBy, sbs.RequirePriceBy, sbs.RequiredWdBy)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
