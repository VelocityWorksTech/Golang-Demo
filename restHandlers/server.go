package restHandlers

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"text/template"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
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

//Start initialise the prerequisites to start the server
func (me *Server) Start() {

	//setup routers.
	me.setupRouters()

	logger := log.New("")

	logger.SetHeader("[${time_rfc3339}] ${level} | ${short_file} | line=${line} |")

	me._client.SetLogger(logger)
	me._e.Logger = logger

	//setup logger middleware
	me._e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	//setup template
	me._e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("./static/index.html")),
	}
	//start the server with graceful shutdown
	me.Run()

}

// Run will run the HTTP Server
func (me *Server) Run() {
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
		if err := me._e.Start(":8000"); err != nil {
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

func (me *Server) setupRouters() {
	dh := NewDataHandler(me._client, me._db)
	me._e.GET("/", dh.Get)
	me._e.GET("/data", dh.GetData)
}
