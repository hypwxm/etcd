FROM golang:1.10.3-alpine as builder

WORKDIR /www/go/src/etcd

COPY * .

RUN CGO_ENABLED=0 GOOS=linux go build etcd

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /www/go/src/etcd/etcd .


CMD ["./etcd"]