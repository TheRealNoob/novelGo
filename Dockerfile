FROM golang:1.24 AS builder

WORKDIR /src

COPY /src /src

RUN go mod download

RUN GOOS=linux go build .

FROM alpine:3.22

USER 1000:1000

COPY --from=builder /src/novelGo /usr/local/bin/novelGo

CMD ["/etc/usrbin/novelGo"]
