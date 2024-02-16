.PHONY: run build run-build tests tests-v tests-coverage tests-coverage-html

run:
	go run -mod=mod cmd/importer/main.go

build:
	go build -mod=mod -o bin/importer cmd/importer/main.go

run-build:
	./bin/importer

tests:
	go test ./...

tests-v:
	go test -coverpkg ./... -v ./...

tests-coverage-html: tests-coverage
	go test -coverpkg ./... -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html

unit-tests:
	go test ./internal...

integration-tests:
	go test -coverpkg ./tests/... -v ./...