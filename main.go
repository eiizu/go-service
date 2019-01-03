package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go-service/controller"
	"go-service/router"
	"go-service/usecase"

	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

func main() {

	logger := logrus.New()
	myRender := render.New()

	myUseCase := usecase.New()
	myController := controller.New(myUseCase, myRender)

	r := router.New(myController)

	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	go func() {
		address := "localhost:7071"
		logger.WithField("Address", address).Info("Starting server")
		if err := http.ListenAndServe(address, r); err != nil {
			logger.Fatal("something went wrong")
		}
	}()

	<-stop

	logger.Info("Shutting down server...")
}
