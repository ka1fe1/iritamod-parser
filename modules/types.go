package msgs

import (
	"github.com/bianjieai/iritamod/modules/identity"
	"github.com/bianjieai/iritamod/modules/node"
	"github.com/bianjieai/iritamod/modules/perm"
	iritaslashing "github.com/bianjieai/iritamod/modules/slashing"
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	models "github.com/kaifei-bianjie/iritamod-parser/types"
)

const (
	MsgTypeUnjail          = "unjail"
	MsgTypeUnjailValidator = "unjail_validator"

	MsgTypeUpdateIdentity = "update_identity"
	MsgTypeCreateIdentity = "create_identity"

	DocTypeAssignRoles   = "assign_roles"
	DocTypeUnassignRoles = "unassign_roles"
	DocTypeBlockAccount  = "block_account"
	DocTypeUnlockAccount = "unblock_account"

	MsgTypeCreateValidator = "create_validator" // type for MsgCreateValidator
	MsgTypeUpdateValidator = "update_validator" // type for MsgUpdateValidator
	MsgTypeRemoveValidator = "remove_validator" // type for MsgRemoveValidator
	MsgTypeGrantNode       = "grant_node"       // type for MsgGrantNode
	MsgTypeRevokeNode      = "revoke_node"      // type for MsgRevokeNode

)

type (
	MsgDocInfo struct {
		DocTxMsg models.TxMsg
		Addrs    []string
		Signers  []string
	}
	SdkMsg sdk.Msg
	Msg    models.Msg

	Coin models.Coin

	Coins []*Coin

	MsgCreateIdentity = identity.MsgCreateIdentity
	MsgUpdateIdentity = identity.MsgUpdateIdentity

	MsgNodeCreate = node.MsgCreateValidator
	MsgNodeUpdate = node.MsgUpdateValidator
	MsgNodeRemove = node.MsgRemoveValidator
	MsgNodeGrant  = node.MsgGrantNode
	MsgNodeRevoke = node.MsgRevokeNode

	MsgAssignRoles    = perm.MsgAssignRoles
	MsgUnassignRoles  = perm.MsgUnassignRoles
	MsgBlockAccount   = perm.MsgBlockAccount
	MsgUnblockAccount = perm.MsgUnblockAccount

	MsgUnjail          = slashing.MsgUnjail
	MsgUnjailValidator = iritaslashing.MsgUnjailValidator
)
