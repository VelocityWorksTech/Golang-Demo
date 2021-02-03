package main

import (
	"html/template"

	"github.com/anil-appface/golang-demo/restHandlers"
	"github.com/anil-appface/golang-demo/store"
	"github.com/anil-appface/golang-demo/utils"
	"github.com/go-resty/resty/v2"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()  //create a new echo
	c := resty.New() //create a new resty

	//initialise db
	db, err := store.OpenDBconnection()
	if err != nil {
		panic(err)
	}
	//setup template
	e.Renderer = &utils.Template{template.Must(template.ParseGlob("static/*.html"))}
	srv := restHandlers.NewServer(e, c, db)

	//start the http server along with dependencies
	srv.Start()
}
