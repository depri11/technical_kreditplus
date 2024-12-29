# Stage 1: Build stage
FROM golang:1.22-alpine AS build

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/api ./cmd

# Stage 2: Final stage
FROM alpine:edge

# Set the working directory
WORKDIR /app

# Set the environment
ENV APP_ENV=development

# Copy the binary from the build stage
COPY --from=build /app/bin/api .

# Copy the config file from the build stage
COPY --from=build /app/config/*.yaml ./config/

# Set the entrypoint command
ENTRYPOINT ["/app/api"]
