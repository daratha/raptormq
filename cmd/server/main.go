package main

import (
	"log"
	"net"
)

func main() {
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

    handleConnection(conn) 
	}
}

func handleConnection(conn net.Conn) {
  defer conn.Close()
  buf := make([]byte, 1024)

  for {
    n, err := conn.Read(buf)
    if err != nil {
      log.Printf("Client %s disconnected", conn.RemoteAddr())
      return
    }
    log.Printf("Message from %s: %q", conn.RemoteAddr(), string(buf[:n]) )
  }
}
