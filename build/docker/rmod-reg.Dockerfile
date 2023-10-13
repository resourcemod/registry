ARG GO_VERSION=1.19

### Build go app
FROM golang:${GO_VERSION} as build
COPY . /rmod-registry
WORKDIR /rmod-registry
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build \
    -mod vendor \
    -o /rmod-reg/bin/rmod-rmod-reg \
    -v /rmod-reg/cmd/rmod-rmod-reg

### Run
FROM alpine:3.7.3
RUN apk update && apk add ca-certificates bash
COPY --from=build /rmod-registry/bin/rmod-registry /usr/local/bin/rmod-registry
ENTRYPOINT ["/bin/bash"]