package layer2

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type DocMsgTransferL2Space struct {
	SpaceId   uint64 `bson:"space_id"`
	Recipient string `bson:"recipient"`
	Sender    string `bson:"sender"`
}

func (m *DocMsgTransferL2Space) GetType() string {
	return MsgTypeTransferL2Space
}

func (m *DocMsgTransferL2Space) BuildMsg(v interface{}) {
	msg := v.(*MsgTransferL2Space)
	m.SpaceId = msg.SpaceId
	m.Recipient = msg.Recipient
	m.Sender = msg.Sender
}

func (m *DocMsgTransferL2Space) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgTransferL2Space)
	addrs = append(addrs, msg.Sender)
	addrs = append(addrs, msg.Recipient)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
