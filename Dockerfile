FROM golang:1.24.1-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/main.go

# expose port the app will run on
EXPOSE 5000

# run the application
CMD ["./main"]
