package main

import (
	"log"
	"sync"
	"time"
)

type hub struct {
	//The mutex to protect connetions
	connectionMx sync.RWMutex

	//Registered connections
	connections map[*connection]struct{}

	//Inbound message connection
	broadcast chan []byte

	logMx sync.RWMutex
	log   [][]byte
}

func newHub() *hub {
	h := &hub{
		connectionMx: sync.RWMutex{},
		broadcast:    make(chan []byte),
		connections:  make(map[*connection]struct{}),
	}

	go func() {
		for {
			msg := <-h.broadcast
			h.connectionMx.Lock()
			for c := range h.connections {
				select {
				case c.send <- msg:
					//stop trying to send to this connection after trying for 1 second.
					//if we have to stop, it means that a reader died so remove the connection also.
				case <-time.After(1 * time.Second):
					log.Printf("shutting down connetion %s", c)
					h.removeConnection(c)
				}
			}
			h.connectionMx.Unlock()
		}
	}()
	return h
}

func (h *hub) addConnection(conn *connection) {
	h.connectionMx.Lock()
	defer h.connectionMx.Unlock()
	h.connections[conn] = struct{}{}
}

func (h *hub) removeConnection(conn *connection) {
	h.connectionMx.Lock()
	defer h.connectionMx.Unlock()
	if _, ok := h.connections[conn]; ok {
		delete(h.connections, conn)
		close(conn.send)
	}
}
