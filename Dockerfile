FROM golang:alpine as builder
LABEL maintiner="saiprasadtoshatwad@gmail.com"
RUN apk update
RUN apk add  git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build cmd/main/main.go -o main
EXPOSE 8080
CMD ["./main"]