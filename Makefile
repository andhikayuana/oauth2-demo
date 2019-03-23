# @author andhikayuana
# @since Jan, 16 2019

GO=$(shell which go)
DEP=$(shell which dep)
DIR=$(shell pwd)
BUILD_DIR=$(DIR)/build
GOARCH=amd64

PROJECT_NAME=oauth2-demo
PACKAGE=github.com/andhikayuana/$(PROJECT_NAME)

install-dep:
	$(DEP) ensure -v

run:
	$(GO) run $(DIR)/cmd/main.go

build: build-linux build-darwin build-windows

build-linux:
	GOOS=linux GOARCH=$(GOARCH) $(GO) build -o $(BUILD_DIR)/$(PROJECT_NAME)_linux_$(GOARCH) $(PACKAGE)/cmd

build-darwin:
	GOOS=darwin GOARCH=$(GOARCH) $(GO) build -o $(BUILD_DIR)/$(PROJECT_NAME)_darwin_$(GOARCH) $(PACKAGE)/cmd

build-windows:
	GOOS=windows GOARCH=$(GOARCH) $(GO) build -o $(BUILD_DIR)/$(PROJECT_NAME)_windows_$(GOARCH).exe $(PACKAGE)/cmd

clean:
	rm -rf build

.PHONY: install-deb run build build-linux build-darwin build-windows