# Builder
FROM golang:1.20.4-alpine3.18 AS builder

LABEL authors="echoedyp"
LABEL company="pondok programmer"
LABEL reachMe="https://github.com/EchoEdyP"

RUN apk update && apk upgrade && apk add --no-cache git curl

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o categoryApp ./app/main.go

# Distribution
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/categoryApp .

EXPOSE 1323

CMD ["./categoryApp"]