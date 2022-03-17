package handler

import "github.com/gin-gonic/gin"

func Upload(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
