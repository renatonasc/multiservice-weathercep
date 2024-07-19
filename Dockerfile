# Estágio de Build
FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/main.go

# Estágio de Publicação
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8000


CMD ["./server"]
