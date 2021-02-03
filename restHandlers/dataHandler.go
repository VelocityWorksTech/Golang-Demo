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

// PageInfoHandler handles the rendering and retrieval of the information to be
// shown.
func (me *dataHandler) Get(c echo.Context) error {
	c.Logger().Debug("Reading from DB")

	var err error
	var catalog *store.Catalog
	url := c.QueryParam("url")
	if url != "" {
		catalog, err = me.saveAndGetData(c, url)
		if err != nil {
			return err
		}
	}

	c.Logger().Info("Rendering template")
	return c.Render(http.StatusOK, "info", catalog)
}

func (me *dataHandler) GetData(c echo.Context) error {
	// c.Echo().Logger
	// app.server.Logger.Info("Reading from DB")
	// info, err := app.fetchAllPayloadRecords()
	// if err != nil {
	// 	return err
	// }

	// page := models.PageData{Payloads: info}

	// app.server.Logger.Info("Rendering template")
	// return c.Render(http.StatusOK, "info", page)
	return nil
}

//performDBaction downlaods data from specific URL and saves it inside db if data is not older than a day
func (me *dataHandler) saveAndGetData(c echo.Context, url string) (*store.Catalog, error) {

	//cleanup if the data for the given url is already exists
	catalog := &store.Catalog{}
	me._db.First(catalog, "URL = ?", url)

	//Delete items & save if the response is not saved
	if time.Now().Sub(catalog.CreatedAt) >= 24*time.Hour {
		datasets := []int64{}
		me._db.Model(&store.Dataset{}).Where(&store.Dataset{}, "CatalogID = ?", catalog.ID).Pluck("ID", &datasets)
		for _, d := range datasets {
			me._db.Delete(&store.Publisher{}).Where("DatasetID=?", d)
			me._db.Delete(&store.Distribution{}).Where("DatasetID=?", d)
			me._db.Delete(&store.ContactPoint{}).Where("DatasetID=?", d)
		}
		me._db.Delete(&store.Dataset{}).Where("CatalogID=?", catalog.ID)
		me._db.Delete(catalog)

		//make request
		c.Logger().Infof("making request url: %s", dataGovURL)
		resp, err := me._client.R().Get(dataGovURL)
		if err != nil {
			return nil, err
		}
		//read response
		catalog = &store.Catalog{}
		c.Logger().Info("parsing the response to catalog")
		catalog.URL = dataGovURL
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
		me._db.Table("datasets").Where("catalog_id = ?", catalog.ID).Find(&catalog.Dataset)
		for i := 0; i < len(catalog.Dataset); i++ {
			me._db.Table("contact_points").Where("dataset_id = ?", catalog.Dataset[i].ID).Find(&catalog.Dataset[i].ContactPoint)
			me._db.Table("publishers").Where("dataset_id = ?", catalog.Dataset[i].ID).Find(&catalog.Dataset[i].Publisher)
			me._db.Table("distributions").Where("dataset_id = ?", catalog.Dataset[i].ID).Find(&catalog.Dataset[i].Distributions)
		}
	}

	return catalog, nil
}
