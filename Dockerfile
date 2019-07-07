FROM golang:1.11.4-alpine as builder
WORKDIR /go/src/github.com/echo-restful-crud-api-example/
COPY . /go/src/github.com/echo-restful-crud-api-example/
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o ./dist/example

FROM alpine:latest
RUN apk add --update ca-certificates
RUN apk add --no-cache tzdata && \
  cp -f /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
  apk del tzdata


WORKDIR /app
COPY ./config/config.yaml /var/app/
COPY ./config/config.yaml /
COPY --from=builder go/src/github.com/echo-restful-crud-api-example/dist/example .

ENV PORT=9090
EXPOSE $PORT
ENTRYPOINT ["./example"]
