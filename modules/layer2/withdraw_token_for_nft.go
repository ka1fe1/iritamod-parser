package layer2

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type DocMsgWithdrawTokenForNFT struct {
	SpaceId uint64 `bson:"space_id"`
	ClassId string `bson:"class_id"`
	TokenId string `bson:"token_id"`
	Owner   string `bson:"owner"`
	Name    string `bson:"name"`
	Uri     string `bson:"uri"`
	UriHash string `bson:"uri_hash"`
	Data    string `bson:"data"`
	Sender  string `bson:"sender"`
}

func (m *DocMsgWithdrawTokenForNFT) GetType() string {
	return MsgTypeWithdrawTokenForNFT
}

func (m *DocMsgWithdrawTokenForNFT) BuildMsg(v interface{}) {
	msg := v.(*MsgWithdrawTokenForNFT)
	m.SpaceId = msg.SpaceId
	m.ClassId = msg.ClassId
	m.TokenId = msg.TokenId
	m.Owner = msg.Owner
	m.Name = msg.Name
	m.Uri = msg.Uri
	m.UriHash = msg.UriHash
	m.Data = msg.Data
	m.Sender = msg.Sender
}

func (m *DocMsgWithdrawTokenForNFT) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgWithdrawTokenForNFT)
	addrs = append(addrs, msg.Sender)
	addrs = append(addrs, msg.Owner)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
