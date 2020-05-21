FROM golang:1.13.7-alpine3.11 AS build
RUN apk update && apk --no-cache add gcc g++ make ca-certificates git
WORKDIR /go/src/github.com/jampajeen/go-cqrs-taxi

COPY . .
RUN go install ./...

FROM alpine:3.11
RUN apk add mysql-client
WORKDIR /usr/bin
COPY --from=build /go/bin .
