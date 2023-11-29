# Dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o main .

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

# Specify the volume
VOLUME ["/bucket"]

EXPOSE 5051
CMD ["./main"]
