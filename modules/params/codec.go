package params

import (
	"github.com/bianjieai/iritamod/modules/params"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(params.AppModuleBasic{})
}
