package msg_parser

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/kaifei-bianjie/iritamod-parser/modules"
)

type Client interface {
	HandleTxMsg(v types.Msg) (msgs.MsgDocInfo, bool)
}
