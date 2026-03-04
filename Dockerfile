FROM docker.io/library/golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 9001
CMD ["./main"]
