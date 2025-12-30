.PHONY: install-libusb build test lint run-sony vendor clean

# Define common variables
GO := go
DOCKER := docker

# Install libusb development headers (Linux/Debian-based systems)
install-libusb:
	sudo apt-get install libusb-1.0-0-dev

build:
	go build ./...

test:
	go test ./... -cover -coverprofile coverage.out -race -mod=vendor -v

lint:
	docker run -t --rm \
	 -v ${PWD}/:/app -w /app \
	 golangci/golangci-lint:v1.56.2 \
	 golangci-lint run -v --config .golangci.yml

# Run the Sony Xperia Z1 Compact brute-force tool
# WARNING: Running 'go run' with 'sudo' is generally discouraged due to security and permission implications.
# Consider using udev rules for USB device access on Linux instead of running the entire program as root.
run-sony:
	sudo go run cmd/sony-xperia-z1-compact/main.go

# Manage Go module dependencies and vendor them
vendor:
	go mod tidy
	go mod vendor

# Clean up build artifacts and coverage files
clean:
	rm -f coverage.out
	rm -rf vendor
	$(GO) clean -cache
