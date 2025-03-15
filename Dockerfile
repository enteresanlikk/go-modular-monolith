FROM golang:1.24

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app/go-modular-monolith

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/go-modular-monolith/app .

CMD ["/app/go-modular-monolith/app"]
