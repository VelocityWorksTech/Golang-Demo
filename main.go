package main

import (
	"github.com/anil-appface/golang-demo/model"
	"github.com/anil-appface/golang-demo/restHandlers"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()  //create a new echo
	c := resty.New() //create a new resty
	db, err := openDBconnection()
	if err != nil {
		panic(err)
	}
	srv := restHandlers.NewServer(e, c, db)

	//start the http server along with dependencies
	srv.Start()
}

//To open & setup db connection
func openDBconnection() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "velocityworks.db")
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.Catalog{}, &model.Distribution{}, &model.Publisher{},
		&model.ContactPoint{}, &model.Dataset{}).Error
	if err != nil {
		return nil, err
	}
	return db, nil
}
