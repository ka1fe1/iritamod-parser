package upgrade

import (
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	"github.com/kaifei-bianjie/iritamod-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		upgrade.AppModuleBasic{},
	)
}
