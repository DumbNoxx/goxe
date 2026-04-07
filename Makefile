# Hi
.PHONY: build dev test dev-args clean

build:
	go build -ldflags="-s -w" -o bin/goxe ./cmd/goxe

dev:
	go run ./cmd/goxe $(ARGS)

test:
	go test -bench=. -benchmem ./cmd/goxe

clean:
	rm -rf ./bin
