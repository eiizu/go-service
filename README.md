# go-service
Template for a REST API service in golang

The service use clean architecture
The dependency injection is in the following order:

Services (any service we need to get data from) or
Store (any database), Should be injected mostly only
to the usecase files.

The usecase uses interfaces to describe the behavior
of the dependencies needed to perform its logic.
Multiple Usecases could use the same dependency.

At this point the dependency injection should look like:
- UseCases1( Service1, Service2, Store1 )
- UseCases2( Service2, Store1 )

The Controllers uses interfaces to describe the behavior
of the dependencies needed to perform its logic.
Controllers mostly only have 1 usecase file.

At this point the dependency injection should look like:
- Controller1( UseCases1( Service1, Service2, Store1 ), config... )
- Controller2( UseCases2( Service2, Store1 ), config... )

We uses interfaces on each layer to describe all the dependencies behavior
In order to do testing we generate mocks from such interfaces

##In order to start the service:
$ go run main.go

##In order to generate mocks after a code change (any interface)
Using mockgen: https://github.com/golang/mock
Install command: $ go install github.com/golang/mock/mockgen@v1.6.0
$ go generate ./...

##In order to execute the unit tests
$ go test ./...

##In order to generate the tests for new functions
Using gotests: https://github.com/cweill/gotests
Install command: $ go get -u github.com/cweill/gotests/...
$ gotests -w -all interal/controllers/<specific-file>.go

##In order to check code with linters
Using golangci-lint: https://golangci-lint.run/usage/linters
Install command: $ go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1
$ golangci-lint run

Using all linters
$ golangci-lint run --enable-all


