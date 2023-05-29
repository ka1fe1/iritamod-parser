package layer2

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type DocMsgUpdateNFTs struct {
	SpaceId uint64        `bson:"space_id"`
	ClassId string        `bson:"class_id"`
	Tokens  []TokenForNFT `bson:"tokens"`
	Sender  string        `bson:"sender"`
}

func (m *DocMsgUpdateNFTs) GetType() string {
	return MsgTypeUpdateNFTs
}

func (m *DocMsgUpdateNFTs) BuildMsg(v interface{}) {
	msg := v.(*MsgUpdateNFTs)
	m.SpaceId = msg.SpaceId
	m.ClassId = msg.ClassId
	m.Sender = msg.Sender
	for _, value := range msg.Tokens {
		m.Tokens = append(m.Tokens, TokenForNFT{
			Id:    value.Id,
			Owner: value.Owner,
		})
	}
}

func (m *DocMsgUpdateNFTs) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgUpdateNFTs)
	addrs = append(addrs, msg.Sender)
	for _, value := range msg.Tokens {
		addrs = append(addrs, value.Owner)
	}
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
