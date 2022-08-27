.PHONY: clean build


vendor: go.mod
	go mod tidy
	go mod vendor

build: vendor
	go build -ldflags="-s -w" -o bin/game main.go

clean:
	go clean -testcache
	rm -f bin/game Gopkg.lock