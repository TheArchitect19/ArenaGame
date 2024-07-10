
FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o main .

FROM alpine:latest

COPY --from=builder /app/main .


EXPOSE 8080

CMD ["./main"]
