test: .test ## run unit tests
.PHONY: .test
.test:
	$(info running tests ...)
	go test ./...


run: .run ## run program
.PHONY: .run
.run:
	$(info running program ...)
	go run cmd/main.go


build: .build ## build an executable
.PHONY: .build
.build:
	$(info building an executable ...)
	go build -o bin/executable cmd/main.go