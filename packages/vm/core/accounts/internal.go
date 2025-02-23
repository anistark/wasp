package accounts

import (
	"errors"

	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/collections"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/vm/core/errors/coreerrors"
)

var (
	ErrNotEnoughFunds                       = coreerrors.Register("not enough funds").Create()
	ErrNotEnoughBaseTokensForStorageDeposit = coreerrors.Register("not enough base tokens for storage deposit").Create()
	ErrNotEnoughAllowance                   = coreerrors.Register("not enough allowance").Create()
	ErrBadAmount                            = coreerrors.Register("bad native asset amount").Create()
	ErrRepeatingFoundrySerialNumber         = coreerrors.Register("repeating serial number of the foundry").Create()
	ErrFoundryNotFound                      = coreerrors.Register("foundry not found").Create()
	ErrOverflow                             = coreerrors.Register("overflow in token arithmetics").Create()
	ErrTooManyNFTsInAllowance               = coreerrors.Register("expected at most 1 NFT in allowance").Create()
	ErrNFTIDNotFound                        = coreerrors.Register("NFTID not found").Create()
)

const (
	// keyAllAccounts stores a map of <agentID> => true
	// where sum = baseTokens + native tokens + nfts
	keyAllAccounts = "a"

	// prefixBaseTokens | <accountID> stores the amount of base tokens (big.Int)
	prefixBaseTokens = "b"
	// prefixBaseTokens | <accountID> stores a map of <nativeTokenID> => big.Int
	prefixNativeTokens = "t"

	// l2TotalsAccount is the special <accountID> storing the total fungible tokens
	// controlled by the chain
	l2TotalsAccount = "*"

	// prefixNFTs | <agentID> stores a map of <NFTID> => true
	prefixNFTs = "n"
	// prefixNFTsByCollection | <agentID> | <collectionID> stores a map of <nftID> => true
	prefixNFTsByCollection = "c"
	// prefixFoundries + <agentID> stores a map of <foundrySN> (uint32) => true
	prefixFoundries = "f"

	// noCollection is the special <collectionID> used for storing NFTs that do not belong in a collection
	noCollection = "-"

	// keyNonce stores a map of <agentID> => nonce (uint64)
	keyNonce = "m"

	// keyNativeTokenOutputMap stores a map of <nativeTokenID> => nativeTokenOutputRec
	keyNativeTokenOutputMap = "TO"
	// keyFoundryOutputRecords stores a map of <foundrySN> => foundryOutputRec
	keyFoundryOutputRecords = "FO"
	// keyNFTOutputRecords stores a map of <NFTID> => NFTOutputRec
	keyNFTOutputRecords = "NO"
	// keyNFTData stores a map of <NFTID> => isc.NFT
	keyNFTData = "ND"
)

func accountKey(agentID isc.AgentID) kv.Key {
	return kv.Key(agentID.Bytes())
}

func allAccountsMap(state kv.KVStore) *collections.Map {
	return collections.NewMap(state, keyAllAccounts)
}

func allAccountsMapR(state kv.KVStoreReader) *collections.ImmutableMap {
	return collections.NewMapReadOnly(state, keyAllAccounts)
}

func accountExists(state kv.KVStoreReader, agentID isc.AgentID) bool {
	return allAccountsMapR(state).HasAt(agentID.Bytes())
}

func allAccountsAsDict(state kv.KVStoreReader) dict.Dict {
	ret := dict.New()
	allAccountsMapR(state).IterateKeys(func(agentID []byte) bool {
		ret.Set(kv.Key(agentID), []byte{0x01})
		return true
	})
	return ret
}

// touchAccount ensures the account is in the list of all accounts
func touchAccount(state kv.KVStore, agentID isc.AgentID) {
	allAccountsMap(state).SetAt([]byte(accountKey(agentID)), codec.EncodeBool(true))
}

// HasEnoughForAllowance checks whether an account has enough balance to cover for the allowance
func HasEnoughForAllowance(state kv.KVStoreReader, agentID isc.AgentID, allowance *isc.Assets) bool {
	if allowance == nil || allowance.IsEmpty() {
		return true
	}
	accountKey := accountKey(agentID)
	if allowance != nil {
		if getBaseTokens(state, accountKey) < allowance.BaseTokens {
			return false
		}
		for _, nativeToken := range allowance.NativeTokens {
			if getNativeTokenAmount(state, accountKey, nativeToken.ID).Cmp(nativeToken.Amount) < 0 {
				return false
			}
		}
	}
	for _, nftID := range allowance.NFTs {
		if !hasNFT(state, agentID, nftID) {
			return false
		}
	}
	return true
}

// MoveBetweenAccounts moves assets between on-chain accounts
func MoveBetweenAccounts(state kv.KVStore, fromAgentID, toAgentID isc.AgentID, assets *isc.Assets) error {
	if fromAgentID.Equals(toAgentID) {
		// no need to move
		return nil
	}

	if !debitFromAccount(state, accountKey(fromAgentID), assets) {
		return errors.New("MoveBetweenAccounts: not enough funds")
	}
	creditToAccount(state, accountKey(toAgentID), assets)

	for _, nftID := range assets.NFTs {
		nft, err := getNFTData(state, nftID)
		if err != nil {
			return err
		}
		if !debitNFTFromAccount(state, fromAgentID, nft) {
			return errors.New("MoveBetweenAccounts: NFT not found in origin account")
		}
		creditNFTToAccount(state, toAgentID, nft)
	}

	touchAccount(state, fromAgentID)
	touchAccount(state, toAgentID)
	return nil
}

func MustMoveBetweenAccounts(state kv.KVStore, fromAgentID, toAgentID isc.AgentID, assets *isc.Assets) {
	err := MoveBetweenAccounts(state, fromAgentID, toAgentID, assets)
	if err != nil {
		panic(err)
	}
}

// debitBaseTokensFromAllowance is used for adjustment of L2 when part of base tokens are taken for storage deposit
// It takes base tokens from allowance to the common account and then removes them from the L2 ledger
func debitBaseTokensFromAllowance(ctx isc.Sandbox, amount uint64) {
	if amount == 0 {
		return
	}
	storageDepositAssets := isc.NewAssetsBaseTokens(amount)
	ctx.TransferAllowedFunds(CommonAccount(), storageDepositAssets)
	DebitFromAccount(ctx.State(), CommonAccount(), storageDepositAssets)
}
