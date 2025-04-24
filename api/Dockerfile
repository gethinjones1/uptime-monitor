# api/Dockerfile

# Use Go on Alpine for a small base
FROM golang:1.23-alpine

# Install Git (for fetching modules) and Air for live reload\
RUN apk add --no-cache git curl \
    && go install github.com/air-verse/air@latest

# Set working directory at project root
WORKDIR /app

# Copy module files for the API and shared modules to cache dependencies
COPY api/go.mod api/go.sum ./api/
COPY shared ./shared

# Download modules (including local replace to shared)
RUN cd api && go mod download

# Copy the rest of the project source
COPY . .

# Switch to the API module directory
WORKDIR /app/api

# Expose port
EXPOSE 8080

# Start using Air for live rebuilds
CMD ["air", "-c", "/app/api/.air.toml"]