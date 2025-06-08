FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env.example .env

RUN go build -o main

FROM alpine:latest

COPY --from=builder /app/main .
COPY --from=builder /app/.env .env

EXPOSE 8080

CMD ["./main"]