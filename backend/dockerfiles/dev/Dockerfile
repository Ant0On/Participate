FROM golang:latest

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod download

CMD ["air", "-c", ".air.toml"]
