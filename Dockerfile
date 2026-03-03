FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o api-server ./cmd/api/main.go

FROM scratch

COPY --from=builder /app/api-server /api-server

EXPOSE 8081

ENTRYPOINT ["/api-server"]