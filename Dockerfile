FROM golang:1.21-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/J3imip/card-validator
COPY . .

RUN go mod tidy
RUN GOOS=linux go build -o /usr/local/bin/card-validator /go/src/github.com/J3imip/card-validator


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/card-validator /usr/local/bin/card-validator
RUN apk add --no-cache ca-certificates

CMD ["sh", "-c", "card-validator"]
