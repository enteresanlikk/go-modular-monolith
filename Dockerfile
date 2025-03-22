FROM golang:1.24.1-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/...

EXPOSE 8080

CMD ["./app"]
