socks-go-runner
===============

Run socks4/socks5 proxy separately with [socks-go](https://github.com/a0s/socks-go).

Options
-------

```shell script
  -host string
        bind to host (default "127.0.0.1")
  -port uint
        bind to port (default 1080)
  -socks4
        enable socks4
  -socks5
        enable socks5
```

Usage
-----

```shell script
go run main.go --host 0.0.0.0 --port 11080 --socks4
```

or

```shell script
docker run --rm -p 11080:11080 a00s/socks-go-runner --host 0.0.0.0 --port 11080 --socks4
```
