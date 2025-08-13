ARG GO_VERSION=1.24.4
FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/pdfium ./cmd/main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/pdfium /app/pdfium
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/config/config.yaml /app/config/config.yaml

CMD ["/app/pdfium"]