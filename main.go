package main

import "github.com/labstack/echo"

func main() {

	//create a new echo
	e := echo.New()
	srv := newServer(e)

	//start the http server along with dependencies
	srv.Start()
}
