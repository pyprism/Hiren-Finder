package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pyprism/Hiren-Finder/controllers"
)

func main() {
	//mapping := bleve.NewIndexMapping()
	//index, err := bleve.New("data", mapping)
	//index, err := bleve.Open("data")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	////data := struct {
	////	Name string
	////}{
	////	Name: "Onion (Red Indian Piyaj) লাল ইন্ডিয়ান পিঁয়াজ",
	////}
	////data2 := struct {
	////	Name string
	////}{
	////	Name: "আলু Potato",
	////}
	////data3 := struct {
	////	Name string
	////}{
	////	Name: "লাল টমেটো Red Tomato",
	////}
	//
	//// index some data
	////err = index.Index("id 1", data)
	////if err != nil {
	////	fmt.Println(err)
	////	return
	////}
	////err = index.Index("id 2", data2)
	////if err != nil {
	////	fmt.Println(err)
	////	return
	////}
	////err = index.Index("id 3", data3)
	////if err != nil {
	////	fmt.Println(err)
	////	return
	////}
	//
	//// search for some text
	//query := bleve.NewFuzzyQuery("আল")
	//query.Fuzziness = 3
	//search := bleve.NewSearchRequest(query)
	//searchResults, err := index.Search(search)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	////fmt.Println(searchResults)
	//for _, v := range searchResults.Hits {
	//	fmt.Println(v.ID)
	//}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/create", controllers.AddSingleIndex)
	r.Run("127.0.0.1:8000") // listen and serve on 0.0.0.0:8080
}
