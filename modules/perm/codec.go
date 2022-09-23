package perm

import (
	"github.com/bianjieai/iritamod/modules/perm"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(perm.AppModuleBasic{})
}
