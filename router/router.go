package router

import (
	"github.com/labstack/echo"
)

// StatusController -
type StatusController interface {
	HandlerStatusz(c echo.Context) error
	HandlerHealthz(c echo.Context) error
}

// SomethingController -
type SomethingController interface {
	HandlerSomething(c echo.Context) error
}

// New -
func New(somethingC SomethingController, statusC StatusController) *echo.Echo {
	e := echo.New()

	e.GET("/statusz", statusC.HandlerStatusz)
	e.GET("/healthz", statusC.HandlerHealthz)

	e.POST("/operation", somethingC.HandlerSomething)

	return e
}
