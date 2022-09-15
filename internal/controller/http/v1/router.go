package v1

import (
	"github.com/gin-gonic/gin"
	"pahan/internal/usecase"
)

func NewRouter(handler *gin.Engine,
	en usecase.Engine,
	su usecase.Suspension) {
	h := handler.Group("/v1")
	{
		newEngineRoutes(h, en)
		newSuspensionRoutes(h, su)
	}
}
