module github.com/kaifei-bianjie/iritamod-parser

go 1.16

require (
	github.com/bianjieai/iritamod v1.2.1-0.20230525103414-2c6d235c2ff1
	github.com/cosmos/cosmos-sdk v0.45.5-0.20220523154235-2921a1c3c918
	github.com/kaifei-bianjie/common-parser v0.0.0-20220923023138-65dfc81a8ff5
	github.com/stretchr/testify v1.7.5
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.45.1-irita-20220816.0.20220816095307-845547d9c19e
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.8-irita-210413.0.20210908054213-781a5fed16d6
)
