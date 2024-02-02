BINARY_NAME=quiz-api

build:
	mkdir -p bin
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME} cmd/quiz-api/main.go

run:
	APP_ENV=production bin/${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm -rf ./bin