FROM golang:1.12.6-alpine3.9 AS build-env

RUN apk add --no-cache git gcc musl-dev
RUN apk add --update make
RUN go get github.com/google/wire/cmd/wire
WORKDIR /go/src/github.com/devtron-labs/external-app-crawler
ADD . /go/src/github.com/devtron-labs/external-app-crawler
RUN GOOS=linux make

FROM alpine:3.9
RUN apk add --no-cache ca-certificates
COPY --from=build-env  /go/src/github.com/devtron-labs/external-app-crawler/external-app-crawler .
CMD ["./external-app-crawler"]