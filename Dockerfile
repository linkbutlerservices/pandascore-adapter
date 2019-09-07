FROM golang:1.12-alpine as builder

ENV GO111MODULE=on

RUN apk add --no-cache make curl git gcc musl-dev linux-headers

WORKDIR /go/src/github.com/linkbutlerservices/pandascore-adapter
ADD . .
RUN make build

# Copy adapter into a second stage container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/linkbutlerservices/pandascore-adapter/pandascore-adapter /usr/local/bin/

EXPOSE 8080
ENTRYPOINT ["pandascore-adapter"]