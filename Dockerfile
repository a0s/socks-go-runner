FROM golang:1.14-alpine AS builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/socks-go-runner

FROM scratch
WORKDIR /app
COPY --from=builder /app/socks-go-runner .
ENTRYPOINT ["/app/socks-go-runner"]
