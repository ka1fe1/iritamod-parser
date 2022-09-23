package perm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type PermClient struct {
}

func NewClient() PermClient {
	return PermClient{}
}

func (nft PermClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgAssignRoles:
		docMsg := DocMsgAssignRoles{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgUnassignRoles:
		docMsg := DocMsgUnassignRoles{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgBlockAccount:
		docMsg := DocMsgBlockAccount{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgUnblockAccount:
		docMsg := DocMsgUnblockAccount{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
