package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/eiizu/go-service/internal/controller"
	"github.com/eiizu/go-service/internal/service"
	"github.com/eiizu/go-service/internal/usecase"

	"github.com/sirupsen/logrus"

	_ "github.com/golang/mock/mockgen/model"
)

const (
	// AppName application name
	AppName = "go-service-demo"
)

func main() {
	logger := logrus.New()

	// Services init
	// Service layer involves all the dependencies that the service
	// needs to perform its operations
	someService := service.NewSomeService("something")

	// UseCase init
	// UseCase layer involves all the business logic of the application,
	// the usecase are split by responsabilities. Eg something, status
	somethingUC := usecase.NewSomething(someService)
	statusUC := usecase.NewStatus(AppName)

	// Controller init
	// Controller layer involves all the http stuff...
	somethingC := controller.NewSomething(somethingUC)
	statusC := controller.NewStatus(statusUC)

	// Create router
	// The router expects multiple controllers with its specific handlers
	r := controller.NewRouter(somethingC, statusC)

	// Define stop signal for the end of excecution
	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	go func() {
		address := ":8080"
		if err := r.Start(address); err != nil {
			logger.Fatal("something went wrong")
		}
	}()

	<-stop

	logger.Info("shutting down server...")
}
