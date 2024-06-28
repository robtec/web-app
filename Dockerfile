FROM golang:1.20-alpine3.19

RUN apk update & apk add curl bash

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o web-app

EXPOSE 8080

CMD ["/app/web-app"]
