package main

import (
	"github.com/anil-appface/golang-demo/restHandlers"
	"github.com/anil-appface/golang-demo/store"
	"github.com/go-resty/resty/v2"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()  //create a new echo
	c := resty.New() //create a new resty
	db, err := store.OpenDBconnection()
	if err != nil {
		panic(err)
	}
	srv := restHandlers.NewServer(e, c, db)

	//start the http server along with dependencies
	srv.Start()
}
