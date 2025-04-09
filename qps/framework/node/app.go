package node

import (
	"encoding/json"
	"fmt"
	"framework/remote"
	"msqp/logs"
)

// App 就是nats的客户端，处理实际游戏逻辑的服务
type App struct {
	remoteCli remote.Client
	readChan  chan []byte
	writeChan chan *remote.Msg
	handlers  LogicHandler
}

func Default() *App {
	return &App{
		readChan:  make(chan []byte, 1024),
		writeChan: make(chan *remote.Msg, 1024),
		handlers:  make(LogicHandler),
	}
}

func (a *App) Run(serverId string) error {
	a.remoteCli = remote.NewNatsClient(serverId, a.readChan)
	err := a.remoteCli.Run()
	if err != nil {
		return err
	}
	go a.readChanMsg()
	go a.writeChanMsg()
	return nil
}

func (a *App) readChanMsg() {
	// 收到的是其他nat client发送的消息
	for {
		select {
		case msg := <-a.readChan:
			var remoteMsg remote.Msg
			fmt.Println("接收到 nat client 发送的消息", msg)
			err := json.Unmarshal(msg, &remoteMsg)
			if err != nil {
				fmt.Println("错误是：", err)
				return
			}
			session := remote.NewSession(a.remoteCli, &remoteMsg)
			session.SetData(remoteMsg.SessionData)
			router := remoteMsg.Router
			//根据路由消息，发送给对应的handler处理
			fmt.Println("router是", router)
			handlerFunc := a.handlers[router]
			if handlerFunc != nil {
				result := handlerFunc(session, remoteMsg.Body.Data)
				message := remoteMsg.Body
				var body []byte
				if result != nil {
					body, _ = json.Marshal(result)
				}
				message.Data = body

				// 得到结果了，发送给connect
				fmt.Println("nat client 发送的消息", remoteMsg.Dst)
				responseMsg := &remote.Msg{
					Src:  remoteMsg.Dst,
					Dst:  remoteMsg.Src,
					Body: message,
					Uid:  remoteMsg.Uid,
					Cid:  remoteMsg.Cid,
				}
				a.writeChan <- responseMsg
			} else {
				fmt.Println("handlerfunc为空")
			}
		}
	}
}

func (a *App) writeChanMsg() {
	for {
		select {
		case msg, ok := <-a.writeChan:
			if ok {
				marshal, _ := json.Marshal(msg)
				err := a.remoteCli.SendMsg(msg.Dst, marshal)
				if err != nil {
					logs.Error("app remote send msg err:%v", err)
				}
			}
		}
	}
}

func (a *App) Close() {
	if a.remoteCli != nil {
		a.remoteCli.Close()
	}
}

func (a *App) RegisterHandler(handler LogicHandler) {
	a.handlers = handler
}
