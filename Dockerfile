FROM golang:1.14-alpine
WORKDIR /app
COPY . /app
RUN go build

FROM golang:1.14-alpine
WORKDIR /app
COPY --from=0 /app/socks-go-runner .
ENTRYPOINT ["/app/socks-go-runner"]
