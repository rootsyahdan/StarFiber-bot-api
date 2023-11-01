FROM golang:1.21.1-alpine3.18 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main

FROM alpine:3.18
WORKDIR /root/
COPY --from=builder /app/main .

COPY .env .

EXPOSE 1312
CMD ["./main"]