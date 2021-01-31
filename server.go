package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo"
)

type server struct {
	_e *echo.Echo
}

func newServer(e *echo.Echo) *server {
	return &server{
		_e: e,
	}
}

func (s *server) Start() {

	//setup routers.

	//start the server with graceful shutdown
	s.Run()

}

// Run will run the HTTP Server
func (s *server) Run() {
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
