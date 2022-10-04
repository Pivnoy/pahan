package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
)

type designRoutes struct {
	t usecase.Model
}

func newModelRoutes(handler *gin.RouterGroup, t usecase.Model) {
	r := &designRoutes{t: t}

	handler.GET("/get_models", r.getAllModels)
	handler.POST("/new_design", r.doNewDesign)
}

type modelResponse struct {
	Models []entity.Model `json:"models"`
}

// GetAllModels godoc
// @Summary list of models
// @Description Get all models
// @Success     200 {array}  entity.Model
// @Failure     400 {object} errResponse
// @Router      /v1/get_models [get]
func (r *designRoutes) getAllModels(c *gin.Context) {
	listModels, err := r.t.GetAllModels(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, modelResponse{listModels})

}

// FIXME
// структура для тела запроса
type doDesignRequest struct {
	WheelDrive   string `json:"wheeldrive"`
	Significance int64  `json:"significance"`
	ProdCost     int64  `json:"prod_cost"`
	EngineID     int64  `json:"engine_id"`
	SuspensionID int64  `json:"suspension_id"`
	VendorID     int64  `json:"vendor_id"`
	Name         string `json:"name"`
}

// FIXME
// проектирование новой модели машины
func (r *designRoutes) doNewDesign(c *gin.Context) {
	var request doDesignRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := r.t.NewModel(c.Request.Context(),
		entity.Model{
			WheelDrive:   request.WheelDrive,
			Significance: request.Significance,
			ProdCost:     request.ProdCost,
			VendorID:     request.VendorID,
			Name:         request.Name,
		})

	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
