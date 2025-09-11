FROM golang:1.24.5-alpine3.22 AS build-stage

ENV GOOS=linux \
    CGO_ENABLED=0 \
    GOARCH=amd64

WORKDIR /app/

COPY go.* /app/

RUN go mod download

COPY . .

RUN go build -o notification-service main.go 

EXPOSE 9002

CMD [ "./notification-service" ]