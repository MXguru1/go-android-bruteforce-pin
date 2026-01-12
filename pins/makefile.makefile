@ -1,12 +1,22 @@
install-libusb:
	sudo apt-get install libusb-1.0-0-dev

build:
	go build -v ./...

test:
	go test ./... -cover -coverprofile coverage.out -race -mod=vendor -v

run:
	go run github.com/jeanbritz/go-android-bruteforce-pin.git/cmd
lint:
	docker run -t --rm \
	 -v ${PWD}/:/app -w /app \
	 golangci/golangci-lint:v1.56.2 \
	 golangci-lint run -v --config .golangci.yml

run-sony:
	sudo go run cmd/sony-xperia-z1-compact/main.go

vendor:
	go mod tidy
	go mod vendor
	go mod vendor
