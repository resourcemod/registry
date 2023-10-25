ARG GO_VERSION=1.19

### Build go app
FROM golang:${GO_VERSION} as build
COPY . /rmod-reg
WORKDIR /rmod-reg
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build -tags netgo -a -v \
    -mod vendor \
    -o /rmod-reg/bin/rmod-reg \
    -v /rmod-reg/cmd/rmod-reg

### Run
FROM alpine:3.9
RUN apk update && apk add ca-certificates bash

COPY --from=build /rmod-reg/bin/rmod-reg /usr/local/bin/rmod-reg