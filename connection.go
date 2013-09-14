package main

import (
	"fmt"
	"github.com/solupro/go.net/websocket"
	"strings"
	"text/template"
	"time"
)

type connection struct {
	// The websocket connection.
	ws *websocket.Conn
	// Buffered channel of outbound messages.
	send        chan string
	name        string
	lastMsgTime time.Time
}

func (c *connection) reader() {
	for {
		var message string
		err := websocket.Message.Receive(c.ws, &message)
		message = template.HTMLEscapeString(message)
		if err != nil {
			break
		}
		if i := strings.Index(message, "name::"); i != -1 {
			name := strings.TrimSpace(message[len("name::"):])

			//check name exist
			for _, v := range clients.memberList() {
				if strings.EqualFold(v, name) || strings.EqualFold("System", name) {
					name = fmt.Sprintf("%s-%d", name, time.Now().Unix()%1000)
					break
				}
			}

			c.name = name
			if q.Len() > 0 {
				c.send <- q.Values.String()
			}
			clients.broadcast <- formatMsg("System", fmt.Sprintf("%s has joined!", name), MSG_WITHMEMBER)

		} else {
			//message = fmt.Sprintf("%s[%s]", message, c.ws.Request().RemoteAddr)
			t := time.Now().In(loc)
			if (int(t.Unix()) - int(c.lastMsgTime.Unix())) < *mf {
				c.send <- formatMsg("System", "you chat frequency too fast!", MSG_WITHMEMBER)
			} else {
				clients.broadcast <- formatMsg(c.name, message, MSG_WITHMEMBER)
				q.Push(formatMsg(c.name, message, MSG_DEFAULT))
				c.lastMsgTime = t
			}
		}
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		err := websocket.Message.Send(c.ws, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}
