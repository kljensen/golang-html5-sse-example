FROM golang:1.9-alpine

WORKDIR /app

RUN apk add --no-cache git \
  && git clone https://github.com/amenezes/golang-html5-sse-example .

EXPOSE 8080

CMD ["run", "./server.go"]
ENTRYPOINT ["/usr/local/go/bin/go"]
