package layer2

import (
	"github.com/bianjieai/iritamod/modules/layer2"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(layer2.AppModuleBasic{})
}
