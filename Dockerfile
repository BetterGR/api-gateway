# syntax=docker/dockerfile:1

# Base image with Go
FROM golang:1.24.3-alpine AS base

# Install build dependencies
RUN apk --no-cache add git

# Set work directory
WORKDIR /app

# =====================
# Dependencies stage
# =====================
FROM base AS deps

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# =====================
# Build stage
# =====================
FROM base AS builder

# Copy dependencies
COPY --from=deps /go/pkg /go/pkg

# Copy the source code
COPY . .

# Build Go binary
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-gateway server.go

# =====================
# Production stage
# =====================
FROM alpine:latest

# Add CA certificates and timezone data
RUN apk --no-cache add ca-certificates tzdata

# Create app directory
WORKDIR /app

# Copy the .env file
COPY --from=builder /app/.env /app/.env

# Set executable binary
COPY --from=builder /api-gateway /app/api-gateway

# Expose the server port (default from .env is 1234)
EXPOSE 1234

# Run the Go application
CMD ["/app/api-gateway"]
