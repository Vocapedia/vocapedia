FROM debian:bullseye as whisper-builder

RUN apt-get update && apt-get install -y build-essential cmake git

WORKDIR /whisper

COPY pkg/whisper ./pkg/whisper

WORKDIR /whisper/pkg/whisper

RUN make

FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY --from=whisper-builder /whisper/pkg/whisper/libwhisper.a ./pkg/whisper/libwhisper.a
COPY --from=whisper-builder /whisper/pkg/whisper/include ./pkg/whisper/include

ENV CGO_ENABLED=1

RUN go build -ldflags="-s -w" -o /app/vocapedia main.go

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y ca-certificates libfftw3-dev bash && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/vocapedia /app/vocapedia
COPY --from=builder /app/config.yaml /app/config.yaml
COPY --from=builder /app/pkg/mail/templates /app/pkg/mail/templates
COPY --from=builder /app/pkg/i18n /app/pkg/i18n
COPY --from=builder /app/pkg/template /app/pkg/template

CMD ["/app/vocapedia"]
