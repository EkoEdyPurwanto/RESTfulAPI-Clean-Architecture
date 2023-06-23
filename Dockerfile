# Builder
FROM golang:1.20.4-alpine3.18 AS builder

LABEL authors="echoedyp"
LABEL company="pondok programmer"
LABEL reachMe="https://github.com/EchoEdyP"

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o categoryApp ./app/main.go

# Distribution
FROM alpine:3.18

WORKDIR /app

RUN apk update && apk upgrade && apk --no-cache add git curl

COPY --from=builder /app/categoryApp .
COPY --from=builder /app/config.json .

EXPOSE 1323

CMD ["./categoryApp"]