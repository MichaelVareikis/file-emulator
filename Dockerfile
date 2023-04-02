FROM golang:1.20.2-alpine3.17 as builder
WORKDIR /app
COPY ./helpers/ ./helpers
COPY go.mod .
COPY main.go .
RUN go build -o app

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/app .
ENTRYPOINT ["./app"]