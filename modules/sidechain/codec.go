package sidechain

import (
	"github.com/bianjieai/iritamod/modules/side-chain"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(sidechain.AppModuleBasic{})
}
