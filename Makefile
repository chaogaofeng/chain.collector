VERSION               := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT                := $(shell git log -1 --format='%H')
DESTDIR         	  ?= bin/chain-exporter

# Build executable file in $DESTDIR directory.
build: go.sum
	@echo "-> Building chain-exporter binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o $(DESTDIR) .

# Install executable file in $GOBIN direcotry. 
install: go.sum
	@echo "-> Installing chain-exporter binary..."
	@go install -mod=readonly $(BUILD_FLAGS) .

# Clean up all generated executable files.
clean:
	@echo "-> Cleaning chain-exporter binary..."
	rm -f $(TOOLS_DESTDIR) 2> /dev/null

PHONY: build install clean