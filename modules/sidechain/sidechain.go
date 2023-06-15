package sidechain

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type SideChainClient struct {
}

func NewClient() SideChainClient {
	return SideChainClient{}
}

func (nft SideChainClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgCreateBlockHeader:
		docMsg := DocMsgCreateBlockHeader{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgCreateSpace:
		docMsg := DocMsgCreateSpace{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgTransferSpace:
		docMsg := DocMsgTransferSpace{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
