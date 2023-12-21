# ==================== builder ====================
FROM golang:1.21.5-alpine as builder
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go run github.com/google/wire/cmd/wire@latest ./...

RUN CGO_ENABLED=0
RUN GOOS=linux
RUN go build cmd/main.go

# ==================== runner ====================
FROM alpine:latest as runner
WORKDIR /app

COPY --from=builder /app/main ./main

EXPOSE 3000

CMD ["./main"]