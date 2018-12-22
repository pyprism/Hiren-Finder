package controllers

import (
	"github.com/blevesearch/bleve"
	"github.com/gin-gonic/gin"
	"github.com/pyprism/Hiren-Finder/conf"
)

func Search(c *gin.Context) {
	sterm := c.PostForm("sterm")

	bleveIndex := conf.Get().BleveIndex
	query := bleve.NewFuzzyQuery(sterm)
	query.Fuzziness = 4
	search := bleve.NewSearchRequest(query)
	searchResults, err := bleveIndex.Search(search)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  err,
		})
	}

	var result []string
	for _, v := range searchResults.Hits {
		result = append(result, v.ID)
	}
	c.JSON(200, gin.H{
		"result":  result,
	})
}
