package v1

import (
	"github.com/gin-gonic/gin"
	"pahan/internal/usecase"
)

func NewRouter(handler *gin.Engine, t usecase.Engine) {
	h := handler.Group("/v1")
	{
		newEngineRoutes(h, t)
	}
}
