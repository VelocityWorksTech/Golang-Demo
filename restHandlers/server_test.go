package restHandlers

import (
	"testing"

	"github.com/labstack/echo"
)

func TestPopulateDB(t *testing.T) {

	e := echo.New()
	s := NewServer(e)
	err := s.PopulateDB()
	if err != nil {
		t.Fail()
	}
}
