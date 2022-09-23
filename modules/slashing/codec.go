package slashing

import (
	iritaslashing "github.com/bianjieai/iritamod/modules/slashing"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(iritaslashing.AppModuleBasic{})
}
