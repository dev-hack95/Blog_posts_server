FROM golang:alpine as builder
LABEL maintiner="saiprasadtoshatwad@gmail.com"
RUN apk update
RUN apk add  git -y
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN cd cmd/main && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main
RUN mv cmd/main/main.sh ./
EXPOSE 8080
CMD ["./main"]