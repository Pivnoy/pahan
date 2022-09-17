package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
)

type designRoutes struct {
	t usecase.Model
}

func newDesignRoutes(handler *gin.RouterGroup, t usecase.Model) {
	r := &designRoutes{t: t}

	handler.POST("/new_design", r.doNewDesign)
}

// структура для тела запроса
type doDesignRequest struct {
	WheelDrive   string  `json:"wheeldrive"`
	Significance int64   `json:"significance"`
	ProdCost     float64 `json:"prod_cost"`
	EngineID     int64   `json:"engine_id"`
	SuspensionID int64   `json:"suspension_id"`
	VendorID     int64   `json:"vendor_id"`
	Name         string  `json:"name"`
}

// проектирование новой модели машины
func (r *designRoutes) doNewDesign(c *gin.Context) {
	var request doDesignRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Fatal(err)
		return
	}

	err := r.t.NewModel(c.Request.Context(),
		entity.Model{
			WheelDrive:   request.WheelDrive,
			Significance: request.Significance,
			ProdCost:     request.ProdCost,
			EngineID:     request.EngineID,
			SuspensionID: request.SuspensionID,
			VendorID:     request.VendorID,
			Name:         request.Name,
		})

	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
