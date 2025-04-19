FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/vocapedia main.go

FROM alpine:latest

RUN apk add --no-cache bash ca-certificates

WORKDIR /app

COPY --from=builder /app/vocapedia /app/vocapedia
COPY --from=builder /app/config.yaml /app/config.yaml
COPY --from=builder /app/pkg/mail/templates /app/pkg/mail/templates
COPY --from=builder /app/pkg/i18n /app/pkg/i18n
COPY --from=builder /app/pkg/template /app/pkg/template

CMD ["/app/vocapedia"]
