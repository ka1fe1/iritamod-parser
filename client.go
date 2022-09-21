package msg_parser

import (
	"github.com/kaifei-bianjie/iritamod-parser/codec"
	"github.com/kaifei-bianjie/iritamod-parser/modules/identity"
	"github.com/kaifei-bianjie/iritamod-parser/modules/params"
	"github.com/kaifei-bianjie/iritamod-parser/modules/perm"
	"github.com/kaifei-bianjie/iritamod-parser/modules/slashing"
	"github.com/kaifei-bianjie/iritamod-parser/modules/upgrade"
)

type MsgClient struct {
	Params   Client
	Slashing Client
	Upgrade  Client
	Identity Client
	Perm     Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Params:   params.NewClient(),
		Slashing: slashing.NewClient(),
		Upgrade:  upgrade.NewClient(),
		Identity: identity.NewClient(),
		Perm:     perm.NewClient(),
	}
}
