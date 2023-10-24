// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package hashdb

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

var GetAccountRoot = getAccountRoot

func getAccountRoot(blob []byte) (common.Hash, error) {
	var account types.StateAccount
	err := rlp.DecodeBytes(blob, &account)
	return account.Root, err
}
