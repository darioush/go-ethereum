// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (bc *BlockContext) Number() *big.Int {
	return bc.BlockNumber
}

func (bc *BlockContext) Timestamp() uint64 {
	return bc.Time
}

func (evm *EVM) ActivePrecompiles() []common.Address {
	if evm.Config.ActivePrecompiles != nil {
		return evm.Config.ActivePrecompiles
	}
	return ActivePrecompiles(evm.chainRules)
}
