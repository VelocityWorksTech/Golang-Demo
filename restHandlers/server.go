package restHandlers

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anil-appface/golang-demo/model"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

const dataGovURL = "https://www.consumerfinance.gov/data.json"

type Server struct {
	_e      *echo.Echo
	_client *resty.Client
	_db     *gorm.DB
}

//NewServer creates new server
func NewServer(e *echo.Echo, c *resty.Client, db *gorm.DB) *Server {
	return &Server{
		_e:      e,
		_client: c,
		_db:     db,
	}
}

func (s *Server) Start() {

	//setup routers.

	//populate db if data doesnt exists in db.
	s.PopulateDB()

	//start the server with graceful shutdown
	s.Run()

}

// Run will run the HTTP Server
func (s *Server) Run() {
	// Set up a channel to listen to for interrupt signals
	var runChan = make(chan os.Signal, 1)

	// Set up a context to allow for graceful server shutdowns in the event
	// of an OS interrupt (defers the cancel just in case)
	_, cancel := context.WithTimeout(
		context.Background(),
		time.Minute*30,
	)
	defer cancel()

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)

	// Run the server on a new goroutine
	go func() {
		if err := s._e.Start(":8000"); err != nil {
			log.Fatalf("Server failed to start due to err: %v", err)
		}
	}()

	// Block on this channel listeninf for those previously defined syscalls assign
	// to variable so we can let the user know why the server is shutting down
	interrupt := <-runChan

	// If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// while alerting the user
	log.Printf("Server is shutting down due to %+v\n", interrupt)
}

//PopulateDB downlaods data from specific URL and saves it inside db
func (s *Server) PopulateDB() error {

	s._e.Logger.Infof("making request url: %s", dataGovURL)
	resp, err := s._client.R().Get(dataGovURL)
	if err != nil {
		return err
	}

	//read response
	s._e.Logger.Info("parsing the response to catalog")
	catalog := &model.Catalog{}
	err = model.ParseCatalogResponse(resp.Body(), catalog)
	if err != nil {
		return err
	}

	//saving to database
	s._e.Logger.Info("Storing to database")
	err = s._db.Create(&catalog).Error
	if err != nil {
		return err
	}

	return nil
}
