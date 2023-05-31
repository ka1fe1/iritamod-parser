package layer2

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type DocMsgCreateL2Space struct {
	Name   string `bson:"name"`
	Uri    string `bson:"uri"`
	Sender string `bson:"sender"`
}

func (m *DocMsgCreateL2Space) GetType() string {
	return MsgTypeCreateL2Space
}

func (m *DocMsgCreateL2Space) BuildMsg(v interface{}) {
	msg := v.(*MsgCreateL2Space)
	m.Name = msg.Name
	m.Uri = msg.Uri
	m.Sender = msg.Sender
}

func (m *DocMsgCreateL2Space) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgCreateL2Space)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
