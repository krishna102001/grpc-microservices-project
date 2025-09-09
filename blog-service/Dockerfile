FROM golang:1.24.5-alpine3.22 AS build-stage

ENV GOOS=linux \
    CGO_ENABLED=0 \
    GOARCH=amd64

WORKDIR /app 

COPY  go.* /app/

RUN go mod download

COPY . .

RUN go build -o blog-service main.go 

FROM alpine:latest

# RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build-stage /app/blog-service .

EXPOSE 9001

CMD [ "./blog-service" ]