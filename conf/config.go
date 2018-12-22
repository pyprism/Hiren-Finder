package conf

import (
	"github.com/blevesearch/bleve"
	"log"
)

type Conf struct {  // bleve global conf
	BleveIndex bleve.Index
}

var config Conf  // shared resource in memory

func Init()  {

	var (
		bleveIndex bleve.Index
		err error = nil
	)

	bleveIndex, err = bleve.Open("data")
	if err != nil {  // create new index file
		mapping := bleve.NewIndexMapping()
		bleveIndex, err = bleve.New( "data", mapping)
		if err != nil {
			log.Fatal(err)
		}
	}

	config = Conf{
		BleveIndex: bleveIndex,
	}

}

// serve single instance of config
func Get() *Conf {
	Init()
	return &config
}
