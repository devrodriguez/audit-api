# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app
COPY . .
RUN go build -o main main.go
EXPOSE 3001
CMD ["/app/main"]