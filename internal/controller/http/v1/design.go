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

	handler.POST("/design", r.doNewDesign)
}

// структура для тела запроса
type doDesignRequest struct {
	ID           int64   `json:"id" binding:"required" example:"1"`
	WheelDrive   string  `json:"wheeldrive" binding:"required" example:"FWD"`
	Significance int64   `json:"significance" binding:"required" example:"1"`
	Price        int64   `json:"price" binding:"required" example:"1000"`
	ProdCost     float64 `json:"prod_cost" binding:"required" example:"1000.1"`
	EngineID     int64   `json:"engine_id" binding:"required" example:"1"`
	SuspensionID int64   `json:"suspension_id" binding:"required" example:"1"`
	VendorID     int64   `json:"vendor_id" binding:"required" example:"1"`
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
			ID:           request.ID,
			WheelDrive:   request.WheelDrive,
			Significance: request.Significance,
			Price:        request.Price,
			ProdCost:     request.ProdCost,
			EngineID:     request.EngineID,
			SuspensionID: request.SuspensionID,
			VendorID:     request.VendorID,
		})

	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
