# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt

test:
	@echo "unit test"
	$(GOTEST) ./...

clean:
	@echo "clean test cache"
	$(GOCLEAN) -testcache

fmt:
	@echo "fmt code"
	$(GOFMT) ./...
