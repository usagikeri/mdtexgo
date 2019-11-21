GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get -u
BINARY_NAME=mt

all: deps build
build:
		# Ubuntu
		GOOS=linux GOARCH=arm $(GOBUILD) -o $(BINARY_NAME) -v mt.go
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
deps:
		$(GOGET) github.com/rakyll/statik
