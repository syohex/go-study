package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	os.Exit(_main())
}

func _main() int {
	port := ":39281"
	if len(os.Args) >= 2 {
		port = ":" + os.Args[1]
	}

	fmt.Printf("Listen TCP server at %s", port)

	listener, err := net.Listen("tcp4", port)
	if err != nil {
		log.Printf("tcp listen error: %v", err)
		return 1
	}
	defer listener.Close()

	for {
		client, err := listener.Accept()
		if err != nil {
			log.Printf("tcp accept error: %v", err)
			return 1
		}

		go handleClient(client)
	}
}

func handleClient(conn net.Conn) {
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.Printf("read client error: %v", err)
			return
		}

		if line == "" {
			return
		}

		upper := strings.ToUpper(line)
		conn.Write([]byte(upper))
	}
}
