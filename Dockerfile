FROM golang:1.13.0-alpine3.10
ENV APP_PATH=/app \
    GOPATH=/golang \
    GOLANG_MIGRATE_VERSION=v4.6.2

ENV DOCKERIZE_VERSION=v0.6.1
RUN wget --quiet https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz && \
    tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz && \
    rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

RUN apk --update add postgresql-client tar curl && \
    curl -sLO https://github.com/golang-migrate/migrate/releases/download/${GOLANG_MIGRATE_VERSION}/migrate.linux-amd64.tar.gz && \
    tar -xzvf migrate.linux-amd64.tar.gz && \
    mv migrate.linux-amd64 /usr/local/bin/migrate
WORKDIR ${APP_PATH}/src
COPY src/go.mod src/go.sum ./
RUN go mod download
