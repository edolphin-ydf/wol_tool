package protocol

import "doll/protocol"

const (
	MsgLoginReq = 1
	MsgLoginRep = 2
	MsgWol      = 3
)

type LoginReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (*LoginReq) GetCmdId() int {
	return MsgLoginReq
}

type LoginRep struct {
	protocol.BaseResponse
}

func (*LoginRep) GetCmdId() int {
	return MsgLoginRep
}

type Wol struct {
	Mac string `json:"mac"`
}

func (*Wol) GetCmdId() int {
	return MsgWol
}


