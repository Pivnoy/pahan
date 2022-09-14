package http

import (
	"github.com/gin-gonic/gin"
	"pahan/internal/usecase"
)

func NewRouter(handler *gin.Engine, t usecase.Engine) {
	h := handler.Group("/test")
	{
		newEngineRoutes(h, t)
	}
}
