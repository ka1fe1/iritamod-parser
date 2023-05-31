package layer2

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type DocMsgDeleteNFTs struct {
	SpaceId  uint64   `bson:"space_id"`
	ClassId  string   `bson:"class_id"`
	TokenIds []string `bson:"token_ids"`
	Sender   string   `bson:"sender"`
}

func (m *DocMsgDeleteNFTs) GetType() string {
	return MsgTypeDeleteNFTs
}

func (m *DocMsgDeleteNFTs) BuildMsg(v interface{}) {
	msg := v.(*MsgDeleteNFTs)
	m.SpaceId = msg.SpaceId
	m.ClassId = msg.ClassId
	m.TokenIds = msg.TokenIds
	m.Sender = msg.Sender
}

func (m *DocMsgDeleteNFTs) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgDeleteNFTs)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
