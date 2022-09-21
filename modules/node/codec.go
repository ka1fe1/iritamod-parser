package node

import (
	"github.com/bianjieai/iritamod/modules/node"
	"github.com/kaifei-bianjie/iritamod-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		node.AppModuleBasic{},
	)
}
