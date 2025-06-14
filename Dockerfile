# Build stage
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# Copy dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# ติดตั้ง tzdata เพื่อรองรับ timezone
RUN apk add --no-cache tzdata

# Copy only the binary from builder
COPY --from=builder /app/main .
COPY .env ./

# Create non-root user
RUN adduser -D appuser
USER appuser

EXPOSE 8080

CMD ["./main"]
