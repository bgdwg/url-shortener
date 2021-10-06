FROM golang:1.17

WORKDIR /go/src/app
COPY . .

RUN go mod tidy
RUN go build -o app ./cmd/main/main.go

CMD ["/go/src/app/app"]