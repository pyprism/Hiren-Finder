package main

import (
	"fmt"
	"github.com/pyprism/Hiren-Finder/config"
	"github.com/pyprism/Hiren-Finder/routes"
)

func main()  {
	config.Init()
	fmt.Println("Running fucking server at http://localhost:4200")
	routes.Init()
}
