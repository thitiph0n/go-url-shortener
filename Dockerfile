FROM golang:alpine as builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app
COPY . .

RUN apk add --no-cache ca-certificates
RUN go mod download
RUN go build -ldflags="-w -s" .

FROM scratch

WORKDIR /app

COPY --from=builder /app/go-url-shortener /usr/bin/

EXPOSE $PORT

ENTRYPOINT ["go-url-shortener"]

