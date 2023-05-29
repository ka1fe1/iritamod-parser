package msgs

import (
	"github.com/bianjieai/iritamod/modules/identity"
	"github.com/bianjieai/iritamod/modules/layer2"
	"github.com/bianjieai/iritamod/modules/node"
	"github.com/bianjieai/iritamod/modules/perm"
	iritaslashing "github.com/bianjieai/iritamod/modules/slashing"
)

const (
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

	MsgTypeCreateL2BlockHeader = "create_l2_block_header"
	MsgTypeCreateL2Space       = "create_l2_space"
	MsgTypeCreateNFTs          = "create_nfts"
	MsgTypeDeleteNFTs          = "delete_nfts"
	MsgTypeDepositClassForNFT  = "deposit_class_for_nft"
	MsgTypeDepositTokenForNFT  = "deposit_token_for_nft"
	MsgTypeTransferL2Space     = "transfer_l2_space"
	MsgTypeUpdateClassesForNFT = "update_classes_for_nft"
	MsgTypeUpdateNFTs          = "update_nfts"
	MsgTypeWithdrawClassForNFT = "withdraw_class_for_nft"
	MsgTypeWithdrawTokenForNFT = "withdraw_token_for_nft"
)

type (
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

	MsgUnjailValidator = iritaslashing.MsgUnjailValidator

	MsgCreateL2BlockHeader = layer2.MsgCreateL2BlockHeader
	MsgCreateL2Space       = layer2.MsgCreateL2Space
	MsgCreateNFTs          = layer2.MsgCreateNFTs
	MsgDeleteNFTs          = layer2.MsgDeleteNFTs
	MsgDepositClassForNFT  = layer2.MsgDepositClassForNFT
	MsgDepositTokenForNFT  = layer2.MsgDepositTokenForNFT
	MsgTransferL2Space     = layer2.MsgTransferL2Space
	MsgUpdateClassesForNFT = layer2.MsgUpdateClassesForNFT
	MsgUpdateNFTs          = layer2.MsgUpdateNFTs
	MsgWithdrawClassForNFT = layer2.MsgWithdrawClassForNFT
	MsgWithdrawTokenForNFT = layer2.MsgWithdrawTokenForNFT
)
