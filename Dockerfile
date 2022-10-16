FROM golang:1.19

WORKDIR /go/src/app

RUN apt-get update

RUN go install github.com/cosmtrek/air@latest && \
    go install github.com/golang/mock/mockgen@v1.6.0

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]
