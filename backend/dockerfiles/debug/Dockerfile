FROM golang:1.21-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .

RUN go build -gcflags="all=-N -l" -o /debug

RUN go mod download

CMD ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/debug"]
