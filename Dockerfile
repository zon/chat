FROM golang:1.24.5-alpine AS build

WORKDIR /usr/src/wurbs

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go env -w GOFLAGS=-buildvcs=false
RUN GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/wurbs-server ./server

FROM alpine:3
COPY --from=build /usr/local/bin/wurbs-server /usr/local/bin/wurbs-server
EXPOSE 8080
CMD ["wurbs-server"]
