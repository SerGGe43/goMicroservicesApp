# Шаг 1: Сборка приложения
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .

# Шаг 2: Финальный контейнер
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8082
CMD ["./main"]
