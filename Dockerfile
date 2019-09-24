FROM golang:1.13.0-alpine3.10
ENV APP_PATH=/app \
    GOPATH=/golang

WORKDIR ${APP_PATH}

COPY go.mod go.sum ${APP_ROOT}/
RUN go mod download
