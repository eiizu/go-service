package controller

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

// StatusUseCase -
type StatusUseCase interface {
	Statusz() (string, error)
	Healthz() (string, error)
}

// Status -
type Status struct {
	UseCase StatusUseCase
	Render  *render.Render
}

// NewStatus -
func NewStatus(uc StatusUseCase, r *render.Render) *Status {
	return &Status{
		UseCase: uc,
		Render:  r,
	}
}

// HandlerStatusz -
func (c *Status) HandlerStatusz(w http.ResponseWriter, r *http.Request) {
	resp, err := c.UseCase.Statusz()
	if err != nil {
		logrus.WithError(err).Info("statusz")
		c.Render.Text(w, http.StatusBadRequest, "something went wrong")
		return
	}

	c.Render.JSON(w, http.StatusOK, resp)
}

// HandlerHealthz -
func (c *Status) HandlerHealthz(w http.ResponseWriter, r *http.Request) {
	resp, err := c.UseCase.Healthz()
	if err != nil {
		logrus.WithError(err).Info("healthz")
		c.Render.Text(w, http.StatusBadRequest, "something went wrong")
		return
	}

	c.Render.JSON(w, http.StatusOK, resp)
}