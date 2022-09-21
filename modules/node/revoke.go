package node

import (
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type DocMsgRevokeNode struct {
	Id       string `bson:"id"`
	Operator string `bson:"operator"`
}

func (m *DocMsgRevokeNode) GetType() string {
	return MsgTypeRevokeNode
}

func (m *DocMsgRevokeNode) BuildMsg(v interface{}) {
	msg := v.(*MsgNodeRevoke)

	m.Id = msg.Id
	m.Operator = msg.Operator
}

func (m *DocMsgRevokeNode) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgNodeRevoke)
	addrs = append(addrs, msg.Operator)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
