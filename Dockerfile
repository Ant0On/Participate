FROM golang:1.18

WORKDIR /participate

COPY . .

RUN go build -o main main.go

CMD ["./main"]