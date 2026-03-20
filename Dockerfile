# Stage 1: Build
FROM docker.io/library/golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Stage 2: Final
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

# Standard-Port definieren, falls nichts von außen kommt
ENV APP_PORT=9001
EXPOSE 9001

CMD ["./main"]
