FROM golang:1.13.0-alpine3.10 as build
ENV GOPATH=/golang
ADD . /build
WORKDIR /build
RUN go build server.go

FROM alpine:3.10
COPY --from=build /build/server /server
EXPOSE 80
CMD ["./server"]
