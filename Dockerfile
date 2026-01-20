FROM golang:1.25-alpine

WORKDIR /app

# Copy dependency list
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build aplikasi
RUN go build -o main ./cmd/api/main.go

# Jalankan aplikasi
CMD ["./main"]