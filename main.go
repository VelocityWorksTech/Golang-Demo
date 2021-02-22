package main

import (
	"github.com/anil-appface/golang-demo/restHandlers"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	//initialising new server
	srv := restHandlers.NewServer()

	//start the http server along with dependencies
	srv.Start()
}
