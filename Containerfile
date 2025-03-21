FROM golang:1.23.7-alpine AS fetch
WORKDIR /usr/src/wurbs
COPY go.mod go.sum ./
RUN go mod download

FROM ghcr.io/a-h/templ:latest AS generate
WORKDIR /usr/src/wurbs
COPY --chown=65532:65532 . .
RUN ["templ", "generate"]

FROM golang:1.23.7-alpine AS build
COPY --from=generate /usr/src/wurbs /usr/src/wurbs
WORKDIR /usr/src/wurbs
RUN go env -w GOFLAGS=-buildvcs=false
RUN go build -o /usr/local/bin/wurbs-server ./server

FROM alpine:3
COPY --from=build /usr/local/bin/wurbs-server /usr/local/bin/wurbs-server
EXPOSE 8080
CMD ["wurbs-server"]