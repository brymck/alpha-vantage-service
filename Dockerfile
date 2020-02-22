# Base build image
FROM golang:alpine as builder

# Install dependencies
RUN apk update && apk add --no-cache ca-certificates git make protobuf-dev tzdata upx && update-ca-certificates
WORKDIR /src

# Use an uncredentialed user
RUN adduser -D -g '' appuser

# Force the Go compiler to use modules
ENV GO111MODULE=on

# Download Go library dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build PROTO_PATH=/usr/include
RUN mv service /bin/service
RUN upx --best /bin/service

# Base deploy image
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

USER appuser

# Expose container
COPY --from=builder /bin/service /bin/service
EXPOSE 8080
ENTRYPOINT ["/bin/service"]
