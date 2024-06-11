// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

func (contract *Contract) AsGenesisContract() *Contract {
	self := AccountRef(contract.Caller())
	if _, ok := contract.caller.(*Contract); ok {
		contract = contract.AsDelegate()
	}
	contract.self = self
	return contract
}
