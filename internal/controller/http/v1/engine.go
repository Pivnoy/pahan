package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pahan/internal/entity"
	"pahan/internal/usecase"
)

type engineRoutes struct {
	t usecase.Engine
}

func newEngineRoutes(handler *gin.RouterGroup, t usecase.Engine) {
	r := engineRoutes{t: t}

	handler.GET("/engine", r.getEngine)
}

type engineResponse struct {
	Engine []entity.Engine `json:"engine"`
}

func (r *engineRoutes) getEngine(c *gin.Context) {
	listEngine, err := r.t.Engines(c.Request.Context())
	if err != nil {
		log.Fatal("engine error")
		return
	}
	c.JSON(http.StatusOK, engineResponse{listEngine})
}
