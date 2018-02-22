package main

import (
	"doll/communicate"
	"net"
	"doll/protocol"
	wolproto "wol_tool/protocol"
	"sync"
	"reflect"
	"goserver/slog"
	"github.com/sabhiram/go-wol"
)

type MsgHandler struct {
	u *User
}

func (h *MsgHandler) HandleLogin(msg *wolproto.LoginReq) {
	if h.u.UserName == msg.UserName && h.u.Password == h.u.Password {
		h.u.isLogin = true
		slog.Trace("登陆成功")
	} else {
		h.u.session.SendMsgSync(&wolproto.LoginRep{protocol.BaseResponse{Code: 1, Msg: "UserName or Password Error"}})
	}
}

func (h *MsgHandler) HandleWol(msg *wolproto.Wol) {
	err := wol.SendMagicPacket(msg.Mac, "255.255.255.255:9", "")
	if err != nil {
		slog.Error(err)
	}
}

var handlerMap = make(communicate.HandlerMap)

func init() {
	communicate.RegisterHandler(&MsgHandler{}, handlerMap)
}

type User struct {
	sync.Mutex
	UserName string
	Password string
	isLogin  bool

	session communicate.Session
}

var handler MsgHandler

func (u *User) MainLoop() {
	for {
		select {
		case msg := <-u.session.GetRecvChan():
			if msg == nil {
				goto End
			}
			communicate.DefaultMsgHandler(handlerMap, msg, []reflect.Value{reflect.ValueOf(&handler)})
			break
		}
	}
End:
	user.session = nil
}

var user = User{
	UserName: "test",
	Password: "test",
}

func main() {
	s := communicate.WebSocketServer{
		HostPort: "localhost:19001",
		UseTLS:   true,
		CertFile: "../wol.crt",
		KeyFile:  "../wol.key",
	}

	handler.u = &user

	s.ListenAndServe(func(conn net.Conn) {
		user.Lock()
		defer user.Unlock()

		if user.session != nil {
			return
		}

		codec := communicate.CreateJsonCodec(handlerMap)
		session := communicate.CreateDefaultChainSession(codec, conn)
		user.session = session

		user.MainLoop()
	})
}
