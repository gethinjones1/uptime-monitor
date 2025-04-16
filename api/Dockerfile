FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# ✅ Build local package, not external
RUN go build -o main .

EXPOSE 8080
CMD ["./main"]
