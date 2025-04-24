package core

import (
  "log"
  "net"
  "sync"
)

type Pubsub struct {
  mu sync.RWMutex
  topics map[string]map[net.Conn]bool
}

func NewPubsub() *Pubsub {
  return &Pubsub {
    topics: make(map[string]map[net.Conn]bool),
  }
}

func (pubsub *Pubsub) Subscribe(conn net.Conn, topic string)  { 
  pubsub.mu.Lock()
  defer pubsub.mu.Unlock()

  if(pubsub.topics[topic] == nil) {
    pubsub.topics[topic] = make(map[net.Conn]bool)
  }

  pubsub.topics[topic][conn] = true
  log.Printf("%s Subscribed to %s", conn.RemoteAddr(), topic)
}

func (pubsub *Pubsub) Unsubscribe(conn net.Conn) {
  pubsub.mu.Lock()
  defer pubsub.mu.Unlock()

  for topic, conns := range pubsub.topics {
    delete(conns, conn)
    if len(conns) == 0 {
      delete(pubsub.topics, topic)
    }
    log.Printf("%s Unsubscribed to %s", conn.RemoteAddr(), topic)
  }
}
