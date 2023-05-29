package layer2

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type Layer2Client struct {
}

func NewClient() Layer2Client {
	return Layer2Client{}
}

func (nft Layer2Client) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgCreateL2BlockHeader:
		docMsg := DocMsgCreateL2BlockHeader{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgCreateL2Space:
		docMsg := DocMsgCreateL2Space{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgTransferL2Space:
		docMsg := DocMsgTransferL2Space{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	//case *MsgCreateNFTs:
	//	docMsg := DocMsgCreateNFTs{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	//	break
	//case *MsgDeleteNFTs:
	//	docMsg := DocMsgDeleteNFTs{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	//	break
	//case *MsgDepositClassForNFT:
	//	docMsg := DocMsgDepositClassForNFT{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	//	break
	//case *MsgDepositTokenForNFT:
	//	docMsg := DocMsgDepositTokenForNft{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	//	break
	//case *MsgUpdateClassesForNFT:
	//	docMsg := DocMsgUpdateClassesForNft{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	//	break
	//case *MsgUpdateNFTs:
	//	docMsg := DocMsgUpdateNFTs{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	//	break
	//case *MsgWithdrawClassForNFT:
	//	docMsg := DocMsgWithdrawClassForNFT{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	//	break
	//case *MsgWithdrawTokenForNFT:
	//	docMsg := DocMsgWithdrawTokenForNFT{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	//	break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
