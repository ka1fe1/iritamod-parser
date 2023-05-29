package layer2

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type TokenForNFT struct {
	Id    string `bson:"id"`
	Owner string `bson:"owner"`
}

type DocMsgCreateNFTs struct {
	SpaceId uint64        `bson:"space_id"`
	ClassId string        `bson:"class_id"`
	Tokens  []TokenForNFT `bson:"tokens"`
	Sender  string        `bson:"sender"`
}

func (m *DocMsgCreateNFTs) GetType() string {
	return MsgTypeCreateNFTs
}

func (m *DocMsgCreateNFTs) BuildMsg(v interface{}) {
	msg := v.(*MsgCreateNFTs)
	m.SpaceId = msg.SpaceId
	m.ClassId = msg.ClassId
	for _, value := range msg.Tokens {
		m.Tokens = append(m.Tokens, TokenForNFT{
			Id:    value.Id,
			Owner: value.Owner,
		})
	}
	m.Sender = msg.Sender
}

func (m *DocMsgCreateNFTs) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgCreateNFTs)
	addrs = append(addrs, msg.Sender)
	for _, value := range msg.Tokens {
		addrs = append(addrs, value.Owner)
	}
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
