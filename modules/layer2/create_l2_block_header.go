package layer2

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type DocMsgCreateL2BlockHeader struct {
	SpaceId uint64 `bson:"space_id"`
	Height  uint64 `bson:"height"`
	Header  string `bson:"header"`
	Sender  string `bson:"sender"`
}

func (m *DocMsgCreateL2BlockHeader) GetType() string {
	return MsgTypeCreateL2BlockHeader
}

func (m *DocMsgCreateL2BlockHeader) BuildMsg(v interface{}) {
	msg := v.(*MsgCreateL2BlockHeader)
	m.SpaceId = msg.SpaceId
	m.Header = msg.Header
	m.Height = msg.Height
	m.Sender = msg.Sender
}

func (m *DocMsgCreateL2BlockHeader) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgCreateL2BlockHeader)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
