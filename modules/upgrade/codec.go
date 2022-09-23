package upgrade

import (
	"github.com/bianjieai/iritamod/modules/upgrade"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		upgrade.AppModuleBasic{},
	)
}
