# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

test:
	@echo "unit test"
	$(GOTEST) ./...

clean:
	@echo "clean test cache"
	$(GOCLEAN) -testcache
