FROM golang:1.19 as build

WORKDIR /sntt

COPY go.mod go.sum ./
COPY Makefile ./main.go ./
COPY ./cmd ./cmd
COPY ./api ./api
COPY ./pkg ./pkg
COPY ./internal ./internal

RUN make build

FROM debian:buster-slim

RUN  apt-get update && apt-get install ca-certificates -y && update-ca-certificates

COPY --from=build /sntt/bin/sntt /usr/local/bin/sntt
ENTRYPOINT ["sntt"]
