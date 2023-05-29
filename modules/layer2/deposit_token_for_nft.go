package layer2

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type DocMsgDepositTokenForNft struct {
	SpaceId uint64 `bson:"space_id"`
	ClassId string `bson:"class_id"`
	TokenId string `bson:"token_id"`
	Sender  string `bson:"sender"`
}

func (m *DocMsgDepositTokenForNft) GetType() string {
	return MsgTypeDepositTokenForNFT
}

func (m *DocMsgDepositTokenForNft) BuildMsg(v interface{}) {
	msg := v.(*MsgDepositTokenForNFT)
	m.SpaceId = msg.SpaceId
	m.ClassId = msg.ClassId
	m.TokenId = msg.TokenId
	m.Sender = msg.Sender
}

func (m *DocMsgDepositTokenForNft) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgDepositTokenForNFT)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
