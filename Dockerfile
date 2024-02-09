FROM golang:1.21.6

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN APP_ENV=production CGO_ENABLED=0 GOOS=linux go build -o ./bin/quiz-api ./cmd/quiz-api/main.go

EXPOSE 8080

CMD ["bin/quiz-api"]