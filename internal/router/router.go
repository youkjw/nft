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

	n := e.Group("/nft")
	e.LoadHTMLGlob("web/*")
	{
		n.GET("/index", handler.Index)
	}
}
