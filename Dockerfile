# Build Stage
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
ENV GIN_MODE=release
RUN go build -o build/main cmd/main.go

# Run Stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/build/main .
ENV GIN_MODE=release
COPY app.env .

EXPOSE 8080
CMD ["/app/main"]
