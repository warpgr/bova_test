FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux \
    go build -a -installsuffix -ldflags="-s -w" \
    -o price-service ./cmd/main.go

FROM alpine:latest

RUN apk update && apk upgrade

RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

RUN adduser -D appuser
USER appuser

WORKDIR /app

COPY --from=builder /app/price-service .

# Set any environment variables required by the application
ENV SERVER_ENDPOINT=:8082
ENV KRAKEN_ENDPOINT="https://api.kraken.com/0/public/Ticker"

EXPOSE 8082

CMD ["./price-service"]
