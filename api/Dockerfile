FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# 👇 Build using Go module-aware path
RUN go build -o main github.com/yourusername/uptime-monitor/api

EXPOSE 8080
CMD ["./main"]
