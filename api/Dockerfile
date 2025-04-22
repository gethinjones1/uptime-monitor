FROM golang:1.23-alpine

WORKDIR /app

# Copy entire project to access go.work & shared/
COPY .. /app

# Set working directory to api module
WORKDIR /app/api

# Download deps in context of the workspace
RUN go mod tidy

# Build
RUN go build -o main .

EXPOSE 8080
CMD ["./main"]
