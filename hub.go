package main

import (
	"fmt"
)

type hub struct {
	// Registered connections.
	connections map[*connection]bool
	// Inbound messages from the connections.
	broadcast chan string
}

func (h *hub) count() int {
	return len(h.connections)
}

func (h *hub) register(c *connection) {
	h.connections[c] = true
}

func (h *hub) unregister(c *connection) {
	delete(h.connections, c)
	close(c.send)
	h.broadcast <- formatMsg("System", fmt.Sprintf("%s has quit!", c.name), MSG_WITHMEMBER)
}

func (h *hub) memberList() (members []string) {
	for conn, _ := range h.connections {
		members = append(members, conn.name)
	}
	return
}

func (h *hub) run() {
	for {
		m := <-h.broadcast
		for c := range h.connections {
			select {
			case c.send <- m:
			default:
				delete(h.connections, c)
				close(c.send)
				go c.ws.Close()
			}
		}
	}
}
