package layer2

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type DocMsgDepositClassForNFT struct {
	SpaceId   uint64 `bson:"space_id"`
	ClassId   string `bson:"class_id"`
	BaseUri   string `bson:"base_uri"`
	Recipient string `bson:"recipient"`
	Sender    string `bson:"sender"`
}

func (m *DocMsgDepositClassForNFT) GetType() string {
	return MsgTypeDepositClassForNFT
}

func (m *DocMsgDepositClassForNFT) BuildMsg(v interface{}) {
	msg := v.(*MsgDepositClassForNFT)
	m.SpaceId = msg.SpaceId
	m.ClassId = msg.ClassId
	m.BaseUri = msg.BaseUri
	m.Recipient = msg.Recipient
	m.Sender = msg.Sender
}

func (m *DocMsgDepositClassForNFT) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgDepositClassForNFT)
	addrs = append(addrs, msg.Sender)
	addrs = append(addrs, msg.Recipient)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
