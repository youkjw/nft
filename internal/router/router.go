package router

import (
	"github.com/gin-gonic/gin"
	"nft/internal/handler"
)

func Router(e *gin.Engine) {
	f := e.Group("/v1/file")
	{
		f.POST("/upload", handler.Upload)
	}
}
