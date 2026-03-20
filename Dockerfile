# Stage 1: Build
FROM docker.io/library/golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Stage 2: Final (Nur die Binärdatei)
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 9001
CMD ["./main"]
