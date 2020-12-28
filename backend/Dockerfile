FROM golang:1.15.6-alpine

RUN apk update && apk add git

RUN go get -u github.com/oxequa/realize && \
    go get -u github.com/gorilla/mux

RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD ./src /go/src/app

CMD ["realize", "start", "--server"]
