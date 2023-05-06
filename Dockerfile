FROM golang:1.20

WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]
