FROM golang:1.23-alpine

WORKDIR /app

# Copy the whole project so we get shared/
COPY .. .

WORKDIR /app/worker

# Ensure dependencies are correct
RUN go mod tidy

RUN go build -o main .

CMD ["./main"]
