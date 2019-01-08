package controller

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

// OperationUseCase -
type OperationUseCase interface {
	Compute(string) (map[string]int, error)
}

// Operation -
type Operation struct {
	UseCase OperationUseCase
	Render  *render.Render
}

// Request -
type Request struct {
	Info string `json:"info"`
}

// New -
func New(uc OperationUseCase, r *render.Render) *Operation {
	return &Operation{
		UseCase: uc,
		Render:  r,
	}
}

// HandlerOperation -
func (op *Operation) HandlerOperation(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data Request
	if err := decoder.Decode(&data); err != nil {
		logrus.WithError(err).Info("decoding")
		op.Render.Text(w, http.StatusBadRequest, "invalid json")
		return
	}

	if data.Info == "" {
		op.Render.Text(w, http.StatusBadRequest, "invalid info")
		return
	}

	resp, err := op.UseCase.Compute(data.Info)
	if err != nil {
		op.Render.Text(w, http.StatusBadRequest, "something went wrong")
		return
	}

	op.Render.JSON(w, http.StatusOK, resp)
}
