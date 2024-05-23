FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd/main/main.go ./

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]