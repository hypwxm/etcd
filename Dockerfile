FROM golang:1.10.3-alpine as builder

# WORKDIR /go/src/etcd/

COPY /var/lib/jenkins/workspace/test/ /go/src/etcd/

RUN cd /go/src/etcd/ && CGO_ENABLED=0 GOOS=linux go build etcd

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /go/src/etcd/etcd .


CMD ["./etcd"]