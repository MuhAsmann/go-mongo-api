# Build stage
FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Final stage (gunakan image sama agar GLIBC kompatibel)
FROM golang:1.24 AS runner

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8000
CMD ["./main"]
