package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// OperationController -
type OperationController interface {
	HandlerOperation(w http.ResponseWriter, r *http.Request)
}

// New -
func New(controller OperationController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.HandlerOperation).Methods(http.MethodGet)

	return r
}
