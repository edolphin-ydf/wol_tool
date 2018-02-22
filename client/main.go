package main

import (
	"doll/communicate"
	"golang.org/x/net/websocket"
	"fmt"
	"crypto/tls"
	"wol_tool/protocol"
	"time"
)

type MsgHandler struct {
}

func (h *MsgHandler) HandleLoginRep(msg *protocol.LoginRep) {
	if msg.Code != 0 {
		fmt.Println(msg.Msg)
	}
}

var handlerMap = make(communicate.HandlerMap)

func init() {
	communicate.RegisterHandler(&MsgHandler{}, handlerMap)
}

func main() {
	codec := communicate.CreateJsonCodec(handlerMap)
	wsCfg, err := websocket.NewConfig("wss://localhost:19001/ws", "wss://localhost:19001/ws")
	if err != nil {
		fmt.Println(err)
		return
	}

	wsCfg.TlsConfig = &tls.Config{InsecureSkipVerify: true}
	cs := communicate.CreateDefaultWSClientWithConfig(codec, wsCfg)
	cs.SendMsg(&protocol.LoginReq{UserName: "test", Password: "test1"})
	cs.SendMsg(&protocol.Wol{Mac:"d8:cb:8a:40:1d:52"})
	time.Sleep(3 * time.Second)
}
