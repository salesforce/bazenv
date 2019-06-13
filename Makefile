#  Copyright (c) 2019, Salesforce.com, Inc.
#  All rights reserved.
#  Licensed under the BSD 3-Clause license.
#  For full license text, see LICENSE.txt file in the repo root  or https://opensource.org/licenses/BSD-3-Clause

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