FROM golang:1.23-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

EXPOSE 50051

ENTRYPOINT ["./binary"]