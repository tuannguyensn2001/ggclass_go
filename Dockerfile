FROM golang:1.18 AS builder
WORKDIR /app
COPY . .
#RUN go build -o main src/server/main.go
EXPOSE 80
CMD ["./main","server"]