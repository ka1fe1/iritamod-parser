package layer2

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type DocMsgWithdrawClassForNFT struct {
	SpaceId uint64 `bson:"space_id"`
	ClassId string `bson:"class_id"`
	Owner   string `bson:"owner"`
	Sender  string `bson:"sender"`
}

func (m *DocMsgWithdrawClassForNFT) GetType() string {
	return MsgTypeWithdrawClassForNFT
}

func (m *DocMsgWithdrawClassForNFT) BuildMsg(v interface{}) {
	msg := v.(*MsgWithdrawClassForNFT)
	m.SpaceId = msg.SpaceId
	m.ClassId = msg.ClassId
	m.Owner = msg.Owner
	m.Sender = msg.Sender
}

func (m *DocMsgWithdrawClassForNFT) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgWithdrawClassForNFT)
	addrs = append(addrs, msg.Sender)
	addrs = append(addrs, msg.Owner)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
