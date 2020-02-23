# Build stage
FROM golang:1.13.8-alpine AS build-env

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/app

# I dont care to copy everything coz its temporary image
COPY . .
RUN go mod download
RUN echo `ls`
RUN cd client && go build && echo `pwd` && echo `ls`

# final stage
FROM alpine:3.9.5
WORKDIR /app
RUN echo `ls`
COPY --from=build-env /go/src/app/client/client client
COPY --from=build-env /go/src/app/env.ini env.ini
ENTRYPOINT ./client
