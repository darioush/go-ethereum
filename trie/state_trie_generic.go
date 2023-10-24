// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package trie

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type StateTrieGeneric[AccountT any] struct {
	*StateTrie
}

func NewStateTrieGeneric[AccountT any](id *ID, db *Database) (*StateTrieGeneric[AccountT], error) {
	trie, err := NewStateTrie(id, db)
	if err != nil {
		return nil, err
	}
	return &StateTrieGeneric[AccountT]{trie}, nil
}

// GetAccount attempts to retrieve an account with provided account address.
// If the specified account is not in the trie, nil will be returned.
// If a trie node is not found in the database, a MissingNodeError is returned.
func (t *StateTrieGeneric[AccountT]) GetAccount(address common.Address) (*AccountT, error) {
	res, err := t.trie.Get(t.hashKey(address.Bytes()))
	if res == nil || err != nil {
		return nil, err
	}
	ret := new(AccountT)
	err = rlp.DecodeBytes(res, ret)
	return ret, err
}

// GetAccountByHash does the same thing as GetAccount, however it expects an
// account hash that is the hash of address. This constitutes an abstraction
// leak, since the client code needs to know the key format.
func (t *StateTrieGeneric[AccountT]) GetAccountByHash(addrHash common.Hash) (*AccountT, error) {
	res, err := t.trie.Get(addrHash.Bytes())
	if res == nil || err != nil {
		return nil, err
	}
	ret := new(AccountT)
	err = rlp.DecodeBytes(res, ret)
	return ret, err
}

// UpdateAccount will abstract the write of an account to the secure trie.
func (t *StateTrieGeneric[AccountT]) UpdateAccount(address common.Address, acc *AccountT) error {
	hk := t.hashKey(address.Bytes())
	data, err := rlp.EncodeToBytes(acc)
	if err != nil {
		return err
	}
	if err := t.trie.Update(hk, data); err != nil {
		return err
	}
	t.getSecKeyCache()[string(hk)] = address.Bytes()
	return nil
}

// Copy returns a copy of StateTrie.
func (t *StateTrieGeneric[AccountT]) Copy() *StateTrieGeneric[AccountT] {
	return &StateTrieGeneric[AccountT]{t.StateTrie.Copy()}
}
