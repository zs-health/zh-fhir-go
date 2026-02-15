# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /zh-fhir ./cmd/zh-fhir

# Final stage
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /zh-fhir .

# Copy the IG data
COPY --from=builder /app/BD-Core-FHIR-IG ./BD-Core-FHIR-IG

# Expose the server port
EXPOSE 8080

# Start the server by default
CMD ["./zh-fhir", "-server", "-port", "8080", "-ig", "./BD-Core-FHIR-IG"]
