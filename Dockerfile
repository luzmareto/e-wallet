# Build Stage
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go
# RUN apk --no-cache add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# Run Stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
# COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
# COPY start.sh .
# COPY wait-for.sh .
# COPY db/migration ./migration

EXPOSE 8080
CMD ["/app/main"]
# ENTRYPOINT [ "/app/start.sh" ]
