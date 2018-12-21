package conf

import (
	"github.com/blevesearch/bleve"
	"log"
	"path/filepath"
	"runtime"
)

type Conf struct {  // bleve global conf
	BleveIndex bleve.Index
}

var config Conf  // shared resource in memory

func Init()  {
	_, b, _, _ := runtime.Caller(0)
	basepath  := filepath.Dir(b)  // get project root directory

	var (
		bleveIndex bleve.Index
		err error
	)

	bleveIndex, err = bleve.Open(basepath + "data")
	if err != nil {  // create new index file
		mapping := bleve.NewIndexMapping()
		bleveIndex, err = bleve.New(basepath + "data", mapping)
		if err != nill {
			log.Fatal(err)
		}
	}

	config = Conf{
		BleveIndex: bleveIndex,
	}

}

// serve single instance of config
func Get() *Conf {
	return &config
}
