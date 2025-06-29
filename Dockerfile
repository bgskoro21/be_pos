FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main

FROM alpine:latest

COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "sh", "-c", "sleep 5 && ./main" ]