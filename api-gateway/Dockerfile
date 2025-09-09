FROM golang:1.24.5-alpine3.22 AS build-stage

ENV GOOS=linux \
    CGO_ENABLED=0 \
    GOARCH=amd64

WORKDIR /app/

COPY go.* .

RUN go mod download

COPY . .

RUN go build -o api-gateway main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=build-stage /app/api-gateway .

EXPOSE 8080

CMD [ "./api-gateway" ]
