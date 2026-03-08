# Build stage
FROM golang:1.25-bookworm AS builder

WORKDIR /src

# Copy dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build a truly static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" -v -o /bin/server .

# Final stage
FROM debian:bookworm-slim

# Install CA certificates
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /bin/server /app/server
RUN chmod +x /app/server

# Copy and prepare the debug entrypoint
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

# Expose port
EXPOSE 8080

# Run the debug entrypoint
ENTRYPOINT ["/app/entrypoint.sh"]
