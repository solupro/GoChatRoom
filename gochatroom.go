package main

import (
	"encoding/json"
	"flag"
	"github.com/solupro/go.net/websocket"
	"log"
	"net/http"
	"text/template"
	"time"
)

const (
	MSG_DEFAULT = iota
	MSG_WITHMEMBER
)

func formatMsg(name, msg string, msgType int) string {
	var m map[string]interface{}
	loc, _ := time.LoadLocation("Asia/Shanghai") //set your timezone
	t := time.Now().In(loc)

	if msgType == MSG_DEFAULT {
		m = map[string]interface{}{
			"from":    name,
			"message": msg,
			"date":    t.Format("2006-01-02 15:04:05"),
		}
	} else if msgType == MSG_WITHMEMBER {
		m = map[string]interface{}{
			"data": map[string]interface{}{
				"from":    name,
				"message": msg,
				"date":    t.Format("2006-01-02 15:04:05"),
			},
			"members": clients.memberList(),
		}
	}
	str, _ := json.Marshal(m)
	return string(str)
}

func wsHandler(ws *websocket.Conn) {
	c := &connection{send: make(chan string, 256), ws: ws, name: ""}
	clients.register(c)
	defer clients.unregister(c)
	go c.writer()
	c.reader()
}

func indexHandler(rw http.ResponseWriter, req *http.Request) {
	tpl, _ := template.ParseFiles("./templates/index.html")
	tpl.Execute(rw, req.Host)
}

var addr = flag.String("addr", ":8177", "http service address")
var Qsize = flag.Int("size", 50, "max size of the message's queue")
var q Queue
var clients = &hub{
	broadcast:   make(chan string),
	connections: make(map[*connection]bool),
}

func main() {
	q.Init(*Qsize)
	flag.Parse()
	go clients.run()

	http.Handle("/ws", websocket.Handler(wsHandler))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", indexHandler)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
