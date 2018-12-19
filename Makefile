# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

test:
	@echo "unit test"
	$(GOTEST) -v ./...

clean:
	@echo "clean test cache"
	$(GOCLEAN) -testcache
