# To build:
# $ docker run --rm -v $(pwd):/go/src/github.com/micahhausler/k8s-signal-logger -w /go/src/github.com/micahhausler/k8s-signal-logger golang:1.7  go build -v -a -tags netgo -installsuffix netgo -ldflags '-w'
# $ docker build -t micahhausler/k8s-signal-logger .
#
# To run:
# $ docker run micahhausler/k8s-signal-logger

FROM alpine

MAINTAINER Micah Hausler, <hausler.m@gmail.com>

RUN apk -U add ca-certificates

COPY k8s-signal-logger /bin/k8s-signal-logger
RUN chmod 755 /bin/k8s-signal-logger

EXPOSE 8080

ENTRYPOINT ["/bin/k8s-signal-logger"]
