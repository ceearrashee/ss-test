FROM golang:latest as builder
WORKDIR /go/src/solid-software
COPY go.mod go.sum ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -installsuffix cgo ./cmd/... && chmod +x /go/bin/* && chmod +x *.sh

FROM alpine:latest
LABEL org.opencontainers.image.authors="Eugene Androsov <eandrosov@gmail.com>"
RUN apk --no-cache add tzdata ca-certificates && \
    rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=builder /go/bin/api-server ./api-server
COPY --from=builder /go/src/solid-software/config.yaml ./config.yaml
EXPOSE 80
CMD ["./api-server"]
