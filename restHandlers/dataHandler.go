package restHandlers

import (
	"net/http"
	"time"

	"github.com/anil-appface/golang-demo/store"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type dataHandler struct {
	_client *resty.Client
	_db     *gorm.DB
}

//NewDataHandler creates the new instance for dataHandler
func NewDataHandler(client *resty.Client, db *gorm.DB) *dataHandler {
	return &dataHandler{
		_client: client,
		_db:     db,
	}
}

//
func (me *dataHandler) Get(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

// Info gets the catalog information & will present in the context
func (me *dataHandler) Info(c echo.Context) error {

	var err error
	catalog := &store.Catalog{}
	url := c.FormValue("urldetails")
	if url != "" {
		catalog, err = me.saveAndGetData(c, url)
		if err != nil {
			return err
		}
	}

	return c.Render(http.StatusOK, "info.html", catalog)

}

func (me *dataHandler) GetData(c echo.Context) error {

	var err error
	catalog := &store.Catalog{}
	url := c.QueryParam("url")
	if url != "" {
		catalog, err = me.saveAndGetData(c, url)
		if err != nil {
			return err
		}
	} else {
		//Get the first record
		catalog.First(me._db)
	}

	return c.JSONPretty(http.StatusOK, catalog, "\t")
}

//performDBaction downlaods data from specific URL and saves it inside db if data is not older than a day
func (me *dataHandler) saveAndGetData(c echo.Context, url string) (*store.Catalog, error) {

	//cleanup if the data for the given url is already exists
	catalog := &store.Catalog{}
	me._db.Where("url = ?", url).First(catalog)

	//Delete items & save if the response is not saved
	if time.Now().Sub(catalog.CreatedAt) >= 24*time.Hour {

		//Delete the catalog
		catalog.Delete(me._db)

		//make request
		c.Logger().Infof("making request url: %s", url)
		resp, err := me._client.R().Get(url)
		if err != nil {
			return nil, err
		}
		//read response
		catalog = &store.Catalog{}
		c.Logger().Info("parsing the response to catalog")
		catalog.URL = url
		err = catalog.Parse(resp.Body())
		if err != nil {
			return nil, err
		}

		//saving to database
		c.Logger().Info("Storing to database")
		if err = me._db.Create(&catalog).Error; err != nil {
			return nil, err
		}
	} else {

		//fetch all datasets
		catalog.GetDatasets(me._db)
	}

	return catalog, nil
}
