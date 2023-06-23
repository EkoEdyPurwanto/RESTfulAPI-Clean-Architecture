FROM golang:alpine

LABEL authors="echoedyp"
LABEL company="pondok programmer"
LABEL reachMe="https://github.com/EchoEdyP"

RUN apk update && apk add --no-cache git curl

WORKDIR /go/src/app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o categoryApp ./app/main.go

CMD ["./categoryApp"]