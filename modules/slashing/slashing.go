package slashing

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type slashingClient struct {
}

func NewClient() slashingClient {
	return slashingClient{}
}

func (slashing slashingClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgUnjail:
		docMsg := DocTxMsgUnjail{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgUnjailValidator:
		docMsg := DocTxMsgUnjailValidator{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
