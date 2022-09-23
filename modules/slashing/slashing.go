package slashing

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type SlashingClient struct {
}

func NewClient() SlashingClient {
	return SlashingClient{}
}

func (slashing SlashingClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgUnjailValidator:
		docMsg := DocTxMsgUnjailValidator{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
