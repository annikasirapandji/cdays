PROJECT?=github.com/annikasirapandji/cdays
BUILD_PATH?=cmd/cdays

build: test
	go build -o ./bin/cdays ./cmd/cdays

test:
	go test -race ./...

