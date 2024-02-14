run:
	go run cmd/importer/main.go

build:
	go build -o bin/importer cmd/importer/main.go

run-build:
	./bin/importer

tests:
	go test ./...

tests-verbose:
	go test -v ./...