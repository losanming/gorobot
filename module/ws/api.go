package ws

import (
	"encoding/json"
	"example.com/m/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type Controller struct{}

// client & serve 的消息体
type msg struct {
	Status int             `json:"status"`
	Data   interface{}     `json:"data"`
	Conn   *websocket.Conn `json:"conn"`
}

// 定义消息类型
const msgTypeOnline = 1  // 上线
const msgTypeOffline = 2 // 离线

var (
	clientMsg = msg{}
	sMsg      = make(chan msg)
	chNotify  = make(chan int, 1)
)

// 注册路由
func (self Controller) RegisterRoute(r *gin.RouterGroup) {
	r.GET("/ws", response.API(Start)) // 后台 - 账号密码登录
}

func Start(c *gin.Context) (data map[string]interface{}, err error) {
	Run(c)
	data = make(map[string]interface{})
	return data, nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Run(gin *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, _ := upgrader.Upgrade(gin.Writer, gin.Request, nil)
	defer ws.Close()

	fmt.Println("Client Connected")
	//开启读写协程

	if err := ws.WriteMessage(1, []byte("Hello client !!! ")); err != nil {
		fmt.Println("err : ", err)
		return
	}
	go reader(ws)
	go writer(ws)
	select {}

}

func reader(conn *websocket.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic err : ", err)
			return
		}
	}()

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		var recv = string(p)
		clientMsg.Data = recv
		sMsg <- clientMsg
		//if err := conn.WriteMessage(messageType, p); err != nil {
		//	fmt.Println("err : ", err)
		//	return
		//}
	}
}

func writer(conn *websocket.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic err : ", err)
			return
		}
	}()
	for {
		select {
		case cl := <-sMsg:
			serveMsgStr, _ := json.Marshal(cl)
			notify(conn, string(serveMsgStr))
		}
	}
}

// 统一消息发放
func notify(conn *websocket.Conn, msg string) {
	chNotify <- 1 // 利用channel阻塞 避免并发去对同一个连接发送消息出现panic: concurrent write to websocket connection这样的异常
	err := conn.WriteMessage(1, []byte(msg))
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	<-chNotify
}
