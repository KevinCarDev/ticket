# Etapa de compilación
FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ticket-app .

# Etapa final: solo el binario
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/ticket-app .

EXPOSE 8080

CMD ["./ticket-app"]