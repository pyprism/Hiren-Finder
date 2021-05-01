package controllers

import (
	"github.com/expectedsh/go-sonic/sonic"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchController struct {}

type SearchForm struct {
	Collection string `json:"collection" binding:"required"`
	Bucket string `json:"bucket" binding:"required"`
	Keywords string `json:"keywords" binding:"required"`
}

func (p *SearchController) PostSearch(c *gin.Context) {
	var searchForm SearchForm

	if err := c.ShouldBindJSON(&searchForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	search, err := sonic.NewSearch("localhost", 1491, "")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	results, error_ := search.Query(searchForm.Collection, searchForm.Bucket, searchForm.Keywords, 100, 0, sonic.LangAutoDetect)
	if error_ != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"search_error": error_.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}