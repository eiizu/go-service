package controller

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

// SomethingUseCase -
type SomethingUseCase interface {
	DoSomething(string) (map[string]int, error)
}

// Something -
type Something struct {
	UseCase SomethingUseCase
	Render  *render.Render
}

// Request -
type Request struct {
	Info string `json:"info"`
}

// NewSomething -
func NewSomething(uc SomethingUseCase, r *render.Render) *Something {
	return &Something{
		UseCase: uc,
		Render:  r,
	}
}

// HandlerSomething -
func (op *Something) HandlerSomething(w http.ResponseWriter, r *http.Request) {
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

	resp, err := op.UseCase.DoSomething(data.Info)
	if err != nil {
		op.Render.Text(w, http.StatusBadRequest, "something went wrong")
		return
	}

	op.Render.JSON(w, http.StatusOK, resp)
}
