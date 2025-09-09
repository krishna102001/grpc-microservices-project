FROM golang:1.24.5-alpine3.22 AS build-stage

ENV GOOS=linux \
    CGO_ENABLED=0 \
    GOARCH=amd64

WORKDIR /app

COPY go.* /app/

RUN go mod download

COPY . .

RUN go build -o auth-service main.go

EXPOSE 9000

CMD [ "./auth-service" ]

# FROM alpine:latest

# WORKDIR /root/

# COPY --from=build-stage /app/auth-service .

# EXPOSE 9000

# CMD [ "./auth-service" ]
