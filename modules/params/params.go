package params

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type paramsClient struct {
}

func NewClient() paramsClient {
	return paramsClient{}
}

func (params paramsClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := false

	return msgDocInfo, ok
}
