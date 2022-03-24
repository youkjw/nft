package error

import "github.com/gin-gonic/gin"

func ParamError(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"data":    nil,
		"message": "Param Error",
		"code":    400,
	})
}

func Error(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"data":    nil,
		"message": "Server Error",
		"code":    400,
	})
}
