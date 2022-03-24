package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	throw "nft/internal/error"
	"nft/internal/utils"
)

type UploadReq struct {
	Name        string                `form:"name" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Price       int64                 `form:"price" binding:"required"`
	File        *multipart.FileHeader `form:"file" binding:"required"`
}

func Upload(c *gin.Context) {
	var req UploadReq
	if err := c.ShouldBind(&req); err != nil {
		throw.ParamError(c, err)
		return
	}

	_, err := req.File.Open()
	if err != nil {
		throw.Error(c, err)
		return
	}

	imageHash, err := utils.Store(c)
	if err != nil {
		throw.Error(c, err)
		return
	}

	fmt.Println(imageHash)

	c.JSON(200, gin.H{
		"data":    nil,
		"message": "",
		"code":    0,
	})
}
