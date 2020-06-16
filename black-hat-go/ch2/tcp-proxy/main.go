package main

import (
	"io"
	"log"
	"net"
)

var (
	target = "news.ycombinator.com:443"
	proxy  = "localhost:1234"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", target)
	if err != nil {
		log.Fatalf("unable to connect to %s\n", target)
	}
	defer dst.Close()

	// Run in goroutine to prevent io.Copy from blocking
	go func() {
		// Copy our source's output to the destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	// Copy our destination's output back to our source
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", proxy)
	log.Printf("listening for connections on %s\n", proxy)
	if err != nil {
		log.Fatalf("unable to bind to %s\n", proxy)
	}

	for {
		conn, err := listener.Accept()
		// log.Printf("connection event")
		if err != nil {
			log.Fatalln("unable to accept connection")
		}
		go handle(conn)
	}
}
