FROM golang:1.22.0-alpine3.17

RUN mkdir /app
COPY .. /app
WORKDIR /app

RUN go build -o api

CMD ["./api"]