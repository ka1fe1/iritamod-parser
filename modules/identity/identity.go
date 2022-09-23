package identity

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type IdentityClient struct {
}

func NewClient() IdentityClient {
	return IdentityClient{}
}

func (nft IdentityClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgCreateIdentity:
		docMsg := DocMsgCreateIdentity{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgUpdateIdentity:
		docMsg := DocMsgUpdateIdentity{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
