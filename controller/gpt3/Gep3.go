package gpt3

import (
	"KpChatGpt/services"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

const month = 60 * 60 * 24 * 30

type SendMsg struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type ReplyMsg struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}

type Client struct {
	ID     string
	sendID string
	Model  string
	Socket *websocket.Conn
	Send   chan []byte
}

type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}

type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var Manager = ClientManager{
	Clients:    make(map[string]*Client), //
	Broadcast:  make(chan *Broadcast),
	Register:   make(chan *Client),
	Reply:      make(chan *Client),
	Unregister: make(chan *Client),
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		c.Socket.PingHandler()
		sendMsg := new(SendMsg)
		err := c.Socket.ReadJSON(&sendMsg)
		if err != nil {
			fmt.Println("数据格式不正确", err)
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}
		if sendMsg.Type == 1 { //如歌tape=1 则为发送消息
			//缓存里面
			//缓存里面
		}

		//cache.RedisClient.incr(c.ID)
		//_, _ = cache.RedisClient.Expire(c.ID, time.Hour*24*30*3).Result()
		//防止过快"分手"  连接过期3个月
		Manager.Broadcast <- &Broadcast{
			Client:  c,
			Message: []byte(sendMsg.Content), //发送过来的消息
			Type:    1,
		}
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			replyMsg := ReplyMsg{
				Code:    200,
				Content: fmt.Sprintf("%s", string(message)),
			}
			msg, _ := json.Marshal(replyMsg)

			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)

		}
	}
}

func (c *Client) ChatWrite() {
	ch1 := make(chan string, 100)
	defer func() {
		_ = c.Socket.Close()
	}()
	go func(ch chan string) {
		for v := range ch {
			replyMsg := ReplyMsg{
				Code:    200,
				Content: fmt.Sprintf("%s", v),
			}
			msg, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}(ch1)
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			services.GetAnswer(string(message), c.Model, ch1)
			//for v := range ch1 {
			//	replyMsg := ReplyMsg{
			//		Code:    200,
			//		Content: fmt.Sprintf("%s", v),
			//	}
			//	msg, _ := json.Marshal(replyMsg)
			//	_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
			//}
		}
	}
}

func (c *Client) Chat() {
	ch := make(chan string, 100)
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		c.Socket.PingHandler()
		sendMsg := new(SendMsg)
		err := c.Socket.ReadJSON(&sendMsg)
		if err != nil {
			fmt.Println("数据格式不正确", err)
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}
		if sendMsg.Type == 2 { //如果tape=1 则为发送消息
			services.GetAnswer(sendMsg.Content, c.Model, ch)
			for v := range ch {
				replyMsg := ReplyMsg{
					Code:    200,
					Content: fmt.Sprintf("%s", v),
				}
				msg, _ := json.Marshal(replyMsg)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
			}
		}
	}
}

func CreateID(uid, toUid string) string {
	return uid + "->" + toUid // 1->2 ()
}

func Gpt(c *gin.Context) {
	uid := c.Query("uid")
	model := c.Query("model")
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}).Upgrade(c.Writer, c.Request, nil) //升级 ws协议
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	//创建一个用户实例
	client := &Client{
		ID:     CreateID(uid, "chat"),
		sendID: CreateID("chat", uid),
		Socket: conn,
		Model:  model,
		Send:   make(chan []byte),
	}
	//创建一个用户实例
	chatClient := &Client{
		ID:     CreateID("chat", uid),
		sendID: CreateID(uid, "chat"),
		Socket: conn,
		Model:  model,
		Send:   make(chan []byte),
	}
	Manager.Register <- chatClient
	//注册到用户管理
	Manager.Register <- client
	go client.Read()
	go client.Write()
	go chatClient.ChatWrite()
	go chatClient.Chat()
}
