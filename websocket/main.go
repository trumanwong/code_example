package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Client struct {
	ID            string          // 连接ID
	Socket        *websocket.Conn // 连接
	HeartbeatTime int64           // 前一次心跳时间
}

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	router := gin.Default()
	router.GET("/ws", func(ctx *gin.Context) {
		conn, err := upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Failed to read message: ", err)
				continue
			}
			log.Printf("Received message: %s\n", msg)
			// 发送消息给客户端
			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Failed to write message:", err)
				break
			}
		}
	})
	router.GET("/client/:client_id", func(ctx *gin.Context) {
		WsChat(ctx.Writer, ctx.Request, ctx.Param("client_id"))
	})
	// Parse Static files
	router.StaticFile("/", "./public/index.html")
	router.Run(":8080")
}
