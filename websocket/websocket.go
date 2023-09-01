package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	message = make(map[string]chan interface{})
	client  = make(map[string]*websocket.Conn)
	mux     sync.Mutex
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WsChat 处理ws请求
func WsChat(w http.ResponseWriter, r *http.Request, clientId string) {
	pingTicker := time.NewTicker(time.Second * 10)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %v\n", err)
		return
	}
	defer conn.Close()
	addClient(clientId, conn)
	m, exist := getMessageChannel(clientId)
	if !exist {
		m = make(chan interface{})
		addMessageChannel(clientId, m)
	}

	conn.SetCloseHandler(func(code int, text string) error {
		deleteClient(clientId)
		fmt.Println(code)
		return nil
	})

	for {
		select {
		case content, ok := <-m:
			err = conn.WriteJSON(content)
			if err != nil {
				log.Println(err)
				if ok {
					go func() {
						m <- content
					}()
				}

				conn.Close()
				deleteClient(clientId)
				return
			}
		case <-pingTicker.C:
			conn.SetWriteDeadline(time.Now().Add(time.Second * 20))
			if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("send ping err:", err)
				conn.Close()
				deleteClient(clientId)
				return
			}
		}
	}
}

func messageHandle(ctx *gin.Context) {
	id := ctx.Param("id")
	if id != "" {
		_, exist := getMessageChannel(id)
		if !exist {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("not exist this client %s", id),
			})
			return
		}
	}

	var m interface{}

	if err := ctx.BindJSON(&m); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "message set failed",
		})
		return
	}

	setMessage(id, m)
}

func addClient(id string, conn *websocket.Conn) {
	mux.Lock()
	client[id] = conn
	mux.Unlock()
}

func getClient(id string) (conn *websocket.Conn, exist bool) {
	mux.Lock()
	conn, exist = client[id]
	mux.Unlock()
	return
}

func deleteClient(id string) {
	mux.Lock()
	delete(client, id)
	mux.Unlock()
}

func addMessageChannel(id string, m chan interface{}) {
	mux.Lock()
	message[id] = m
	mux.Unlock()
}

func getMessageChannel(id string) (m chan interface{}, exist bool) {
	mux.Lock()
	m, exist = message[id]
	mux.Unlock()
	return
}

func setMessage(id string, content interface{}) {
	mux.Lock()
	if m, exist := message[id]; exist {
		go func() {
			m <- content
		}()
	}
	mux.Unlock()
}

func setMessageAllClient(content interface{}) {
	mux.Lock()
	all := message
	mux.Unlock()
	go func() {
		for _, m := range all {
			m <- content
		}
	}()

}

func deleteMessageChannel(id string) {
	mux.Lock()
	if m, ok := message[id]; ok {
		close(m)
		delete(message, id)
	}
	mux.Unlock()
}
