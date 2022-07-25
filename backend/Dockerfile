# BASE
FROM golang:1.18.4-alpine as base

RUN apk update && apk add --no-cache git gcc musl-dev

RUN mkdir /go/src/studyapp
WORKDIR /go/src/studyapp

COPY src/go.mod .
COPY src/go.sum .
RUN go mod download
RUN go install github.com/cosmtrek/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/rubenv/sql-migrate/...@latest

COPY ./src /go/src/studyapp

CMD ["air", "-c", ".air.toml"]

# BUILDER
FROM golang:1.18.4-alpine as builder

RUN mkdir /go/src/builder
WORKDIR /go/src/builder

COPY --from=base /go/src/studyapp /go/src/builder
RUN CGO_ENABLED=0 go build ./main.go

# PROD
FROM alpine as production

RUN mkdir -p /go/src
WORKDIR /go/src

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/builder/main /go/src/main

EXPOSE 8080

CMD [ "/go/src/main" ]
