package restHandlers

import "github.com/labstack/echo"

type dataHandler struct {
}

//NewDataHandler creates the new instance for dataHandler
func NewDataHandler() *dataHandler {
	return &dataHandler{}
}

// PageInfoHandler handles the rendering and retrieval of the information to be
// shown.
func (me *dataHandler) Get(c echo.Context) error {
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

func (me *dataHandler) postHandler(c echo.Context) error {

	return nil
}
