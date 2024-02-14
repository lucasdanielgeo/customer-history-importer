run:
	go run -mod=mod cmd/importer/main.go

build:
	go build -mod=mod -o bin/importer cmd/importer/main.go

run-build:
	./bin/importer

tests:
	go test ./...

tests-v:
	go test -v ./...