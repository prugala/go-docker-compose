ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine

RUN apk add --no-cache git
RUN ${GO_VERSION}
WORKDIR /app