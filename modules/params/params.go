package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
)

type ParamsClient struct {
}

func NewClient() ParamsClient {
	return ParamsClient{}
}

func (params ParamsClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	return msgDocInfo, false
}
