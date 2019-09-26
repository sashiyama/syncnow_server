FROM golang:1.13.0-alpine3.10
ENV APP_PATH=/app \
    GOPATH=/golang \
    GOLANG_MIGRATE_VERSION=v4.6.2
RUN apk --update add postgresql-client tar curl && \
    curl -sLO https://github.com/golang-migrate/migrate/releases/download/${GOLANG_MIGRATE_VERSION}/migrate.linux-amd64.tar.gz && \
    tar -xzvf migrate.linux-amd64.tar.gz && \
    mv migrate.linux-amd64 /usr/local/bin/migrate
WORKDIR ${APP_PATH}/src
COPY src/go.mod src/go.sum ./
RUN go mod download
