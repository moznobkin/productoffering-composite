FROM golang:1.12.7-alpine3.10 AS builder
RUN apk add git

RUN mkdir -p /go/src/github.com/moznobkin/productoffering-composite
COPY . /go/src/github.com/moznobkin/productoffering-composite
WORKDIR /go/src/github.com/moznobkin/productoffering-composite
ENV GO111MODULE=on
RUN go build -v -o /go/bin/productoffering-composite.exe

FROM alpine:3.10
LABEL maintainer="Michael Oznobkin <oznobkin@gmail.com>"

COPY --from=builder /go/bin/productoffering-composite.exe /usr/bin/productoffering-composite.exe
COPY --from=builder /go/src/github.com/moznobkin/productoffering-composite/examples/json/  /usr/bin/examples/json/

WORKDIR /usr/bin
ARG Q_PASS

EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/productoffering-composite.exe"]