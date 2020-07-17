package main

import (
	"flag"
	"fmt"
	"github.com/a0s/socks-go"
	"log"
	"net"
	"strings"
	"time"
)

var (
	host          = flag.String("host", "127.0.0.1", "ip address to bind")
	port          = flag.Uint64("port", 1080, "port number to bind")
	socks4Enabled = flag.Bool("socks4", false, "enable socks4")
	socks5Enabled = flag.Bool("socks5", false, "enable socks5")
)

func main() {
	flag.Parse()

	var protocols []string
	if *socks4Enabled == true {
		protocols = append(protocols, "socks4")
	}
	if *socks5Enabled == true {
		protocols = append(protocols, "socks5")
	}
	if len(protocols) == 0 {
		protocols = append(protocols, "none")
	}
	withProtocols := strings.Join(protocols, ",")
	address := fmt.Sprintf("%v:%v", *host, *port)
	log.Printf("bind to %v with %v", address, withProtocols)

	conn, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := conn.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("connected from %s", c.RemoteAddr())

		d := net.Dialer{Timeout: 10 * time.Second}
		s := socks.Conn{Conn: c, Dial: d.Dial, Socks4Enabled: *socks4Enabled, Socks5Enabled: *socks5Enabled}
		go s.Serve()
	}
}
