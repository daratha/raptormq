package main

import (
  "github.com/daratha/raptormq/internal/core"
	"log"
	"net"
)

func main() {
  pubsub := core.NewPubsub()
	// Start TCP server on port 6000
	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	log.Println("TCP server listening on :6000")

	// Accept connections forever
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

    go handleConnection(conn, pubsub) 
	}
}

func handleConnection(conn net.Conn, pubsub *core.Pubsub) {
  defer conn.Close()
  buf := make([]byte, 1024)

  for {
    n, err := conn.Read(buf)
    if err != nil {
      log.Printf("Client %s disconnected", conn.RemoteAddr())
      return
    }
    log.Printf("Message from %s: %q", conn.RemoteAddr(), string(buf[:n]) )

    if string(buf[:3]) == "SUB" {
      topic := string(buf[4:n-1])
      pubsub.Subscribe(conn, topic)
    }
  }
}
