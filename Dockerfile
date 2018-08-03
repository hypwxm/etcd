FROM golang:1.10.3-alpine as builder

WORKDIR /go/src/etcd/

COPY / .

RUN go build etcd

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /go/src/etcd/etcd .
COPY --from=0 /go/src/etcd/tls/ /tls/


CMD ["./etcd"]