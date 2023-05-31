package layer2

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type UpdateClassForNFT struct {
	Id    string `bson:"id"`
	Uri   string `bson:"uri"`
	Owner string `bson:"owner"`
}

type DocMsgUpdateClassesForNft struct {
	SpaceId            uint64              `bson:"space_id"`
	ClassUpdatesForNft []UpdateClassForNFT `bson:"class_updates_for_nft"`
	Sender             string              `bson:"sender"`
}

func (m *DocMsgUpdateClassesForNft) GetType() string {
	return MsgTypeUpdateClassesForNFT
}

func (m *DocMsgUpdateClassesForNft) BuildMsg(v interface{}) {
	msg := v.(*MsgUpdateClassesForNFT)
	m.SpaceId = msg.SpaceId
	m.Sender = msg.Sender
	for _, value := range msg.ClassUpdatesForNft {
		m.ClassUpdatesForNft = append(m.ClassUpdatesForNft, UpdateClassForNFT{
			Id:    value.Id,
			Uri:   value.Uri,
			Owner: value.Owner,
		})
	}
}

func (m *DocMsgUpdateClassesForNft) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgUpdateClassesForNFT)
	addrs = append(addrs, msg.Sender)
	for _, value := range msg.ClassUpdatesForNft {
		addrs = append(addrs, value.Owner)
	}
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
