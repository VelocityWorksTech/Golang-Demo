package restHandlers

import (
	"testing"

	"github.com/anil-appface/golang-demo/store"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

func TestPopulateDB(t *testing.T) {

	e := echo.New()  //create a new echo
	c := resty.New() //create a new resty
	db, err := testDBconnection()
	if err != nil {
		panic(err)
	}
	dataHandler := NewDataHandler(c, db)
	catalog, err := dataHandler.saveAndGetData(e.AcquireContext())
	if err != nil {
		t.Fail()
	}

	if catalog.URL == "" {
		t.Fail()
	}
}

//To open & setup db connection
func testDBconnection() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "velocityworks_test.db")
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&store.Catalog{}, &store.Distribution{}, &store.Publisher{},
		&store.ContactPoint{}, &store.Dataset{}).Error
	if err != nil {
		return nil, err
	}
	return db, nil
}
