package iritamod_parser

import (
	common "github.com/kaifei-bianjie/common-parser"
	"github.com/kaifei-bianjie/iritamod-parser/codec"
	"github.com/kaifei-bianjie/iritamod-parser/modules/identity"
	"github.com/kaifei-bianjie/iritamod-parser/modules/node"
	"github.com/kaifei-bianjie/iritamod-parser/modules/params"
	"github.com/kaifei-bianjie/iritamod-parser/modules/perm"
	"github.com/kaifei-bianjie/iritamod-parser/modules/slashing"
	"github.com/kaifei-bianjie/iritamod-parser/modules/upgrade"
)

type MsgClient struct {
	Params   common.Client
	Slashing common.Client
	Upgrade  common.Client
	Identity common.Client
	Perm     common.Client
	Node     common.Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Params:   params.NewClient(),
		Slashing: slashing.NewClient(),
		Upgrade:  upgrade.NewClient(),
		Identity: identity.NewClient(),
		Perm:     perm.NewClient(),
		Node:     node.NewClient(),
	}
}
