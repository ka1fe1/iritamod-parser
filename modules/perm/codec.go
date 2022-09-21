package perm

import (
	"github.com/bianjieai/iritamod/modules/perm"
	"github.com/kaifei-bianjie/iritamod-parser/codec"
)

func init() {
	codec.RegisterAppModules(perm.AppModuleBasic{})
}
