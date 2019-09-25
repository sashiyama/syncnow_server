FROM golang:1.13.0-alpine3.10
ENV APP_PATH=/app \
    GOPATH=/golang

WORKDIR ${APP_PATH}

COPY src/go.mod src/go.sum ./src/

WORKDIR ${APP_PATH}/src

RUN go mod download
