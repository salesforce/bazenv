# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
GONAMESPACE=github.com/salesforce/bazenv

all: test build
build: 		
		$(GOBUILD) $(GONAMESPACE)/cmd/bazenv
		$(GOBUILD) $(GONAMESPACE)/cmd/bazel
test: 
		$(GOTEST) $(GONAMESPACE)/pkg/bazenv
clean: 
		$(GOCLEAN)
		rm -f bazenv
		rm -f bazel
install: 		
		$(GOINSTALL) $(GONAMESPACE)/cmd/bazenv
		$(GOINSTALL) $(GONAMESPACE)/cmd/bazel
deps:
		dep ensure