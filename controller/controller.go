package controller

import (
	"net/http"

	"github.com/unrolled/render"
)

// OperationUseCase -
type OperationUseCase interface {
	Compute() (string, error)
}

// Operation -
type Operation struct {
	UseCase OperationUseCase
	Render  *render.Render
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
	resp, err := op.UseCase.Compute()
	if err != nil {
		op.Render.Text(w, http.StatusBadRequest, "something went wrong")
		return
	}

	op.Render.Text(w, http.StatusOK, resp)
}
