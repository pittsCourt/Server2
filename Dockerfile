# syntax=docker/dockerfile:1
FROM golang:1.18-alpine
WORKDIR /app
COPY vendor /app/vendor
COPY handlers /app/handlers
COPY config.yaml /app/config.yaml
COPY go.mod ./
COPY go.sum ./
RUN apk update && apk add git
COPY *.go ./
RUN go build -o /server2
EXPOSE 80
CMD [ "/server2" ]