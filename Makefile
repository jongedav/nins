CMD=./cmd/nins

.PHONY: run build tidy clean

run:
	go run $(CMD)

build:
	go build -o bin/$(BINARY) $(CMD)

tidy:
	go mod tidy

clean:
	rm -rf bin
