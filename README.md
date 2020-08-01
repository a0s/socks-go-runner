socks-go-runner
===============

Run socks4/socks5 proxy separately with [socks-go](https://github.com/a0s/socks-go).

Options
-------

```shell script
  -host string
        ip address to bind (default "127.0.0.1")
  -port uint
        port number to bind (default 1080)
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
