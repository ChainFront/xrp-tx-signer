package xrptx

import (
	"fmt"
	"github.com/awnumar/memguard"
	"github.com/rubblelabs/ripple/crypto"
	"github.com/rubblelabs/ripple/data"
)

//SignTx Sign a transaction given a private key. Returns a hex encoded signed tx.
func SignTx(txm data.TransactionWithMetaData, seed memguard.LockedBuffer) (*string, error) {

	keySequence := uint32(0)

	key, err := crypto.NewECDSAKey(seed.Buffer())
	defer zeroKey(key)

	// Sign the transaction
	tx := txm.Transaction
	err = data.Sign(tx, key, &keySequence)
	if err != nil {
		return nil, err
	}

	// Convert to hex
	_, txRaw, err := data.Raw(tx)
	if err != nil {
		return nil, err
	}
	encodedTx := fmt.Sprintf("%X", txRaw)

	return &encodedTx, nil
}

//ZeroKey Removes key data from memory
func zeroKey(k crypto.Key) {
	b := k.Private(nil)
	for i := range b {
		b[i] = 0
	}
}
