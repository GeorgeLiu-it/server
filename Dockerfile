# ---- Build stage ----
FROM golang:1.23.2-alpine AS builder

# Add your private CA
COPY Zscaler_Root_CA.crt /usr/local/share/ca-certificates/ca.crt
RUN update-ca-certificates

# Add git + ca-certificates (required for HTTPS + proxy trust)
RUN apk add --no-cache git ca-certificates

# Set working dir
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# ---- Final stage ----
FROM scratch

# Copy binary
COPY --from=builder /app/app /app

# Copy config
COPY config.yaml /config.yaml

# Copy CA certs (including your private CA)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app"]