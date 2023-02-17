FROM golang:alpine AS builder

WORKDIR /app

ADD certs/client.crt /etc/ssl/certs/

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./cmd/twitter-tools


FROM scratch

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/.env ./.env
COPY --from=builder /app/main /usr/bin/

CMD [ "main" ]