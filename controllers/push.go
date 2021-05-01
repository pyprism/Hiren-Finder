package controllers

import (
	"github.com/expectedsh/go-sonic/sonic"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PushForm struct {
	Collection string `json:"collection" binding:"required"`
	Bucket string `json:"bucket" binding:"required"`
	Identifier string `json:"identifier" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type PushController struct {}

func (p *PushController) PostData(c *gin.Context) {

	var pushForm PushForm
	if err := c.ShouldBindJSON(&pushForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ingester, err := sonic.NewIngester("localhost", 1491, "")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ingester.Push( pushForm.Collection, pushForm.Bucket, pushForm.Identifier, pushForm.Message, sonic.LangAutoDetect)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
