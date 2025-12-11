# Go image
FROM golang:1.25 AS builder

WORKDIR /app

# Go mod dosyalarını kopyala
COPY go.mod go.sum ./
RUN go mod download

# Tüm projeyi kopyala
COPY . .
RUN go build -o server .


# Final image
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 7070

CMD ["./server"]
