package identity

import (
	"github.com/bianjieai/iritamod/modules/identity"
	"github.com/kaifei-bianjie/iritamod-parser/codec"
)

func init() {
	codec.RegisterAppModules(identity.AppModuleBasic{})
}
