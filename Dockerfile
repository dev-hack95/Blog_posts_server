FROM golang:alpine as builder
LABEL maintiner="saiprasadtoshatwad@gmail.com"
RUN apk update && apk add --no-achce git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN cd cmd/main && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main
EXPOSE 8080
CMD ["./main"]