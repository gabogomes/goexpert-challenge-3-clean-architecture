FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/ordersystem

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/ordersystem main.go wire_gen.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/ordersystem /app/ordersystem

COPY --from=builder /app/cmd/ordersystem/.env /app/.env

EXPOSE 8080

CMD ["/app/ordersystem"]
