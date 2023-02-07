package handle

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func (manager *ClientManager) Start() {
	for {
		fmt.Println("------监听管道通信------")
		select {
		case conn := <-Manager.Register:
			fmt.Printf("有新链接:%v \n", conn.ID)
			Manager.Clients[conn.ID] = conn //有链接放到用户管理上
			replyMsg := ReplyMsg{
				Code:    200,
				Content: "已经连接到服务器了",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-Manager.Unregister:
			fmt.Printf("连接失败:%v \n", conn.ID)
			if _, ok := Manager.Clients[conn.ID]; ok {
				replyMsg := &ReplyMsg{
					Code:    500,
					Content: "连接中断",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
			}
		case broadcast := <-Manager.Broadcast: //如果是1发送给2
			message := broadcast.Message
			sendId := broadcast.Client.sendID //2 接受 1 的消息
			flag := false                     //默认对方是不在线的
			for id, conn := range Manager.Clients {
				if id != sendId {
					continue
				}
				select {
				case conn.Send <- message:
					flag = true
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
			}
			_ = broadcast.Client.ID
			if flag {
				replyMsg := &ReplyMsg{
					Code:    200,
					Content: "对方在线应答",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				//插入到数据库中--

				//end------
			} else {
				fmt.Println("对方不在线")
				replyMsg := &ReplyMsg{
					Code:    200,
					Content: "对方不在线",
				}
				msg, err := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					fmt.Println("对方不在线应答 Err", err)
				}
			}
		}

	}

}
