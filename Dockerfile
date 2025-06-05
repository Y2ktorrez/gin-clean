FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/gin-clean ./cmd/app

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/gin-clean .
COPY .env .

EXPOSE 8080

CMD ["./gin-clean"]
