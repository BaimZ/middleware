package app

import (
	"fmt"
	"log"
	"middleware/internal/app/endpoint"
	"middleware/internal/app/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(MW)

	a.echo.GET("/status", a.e.Status)

	return a, nil

}

func (a *App) Run() error {
	fmt.Println("server is running")
	err := s.Start(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
