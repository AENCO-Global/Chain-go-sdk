package transaction

import (
	"github.com/slackve/nem2-sdk-go/core/coders"
)

// Transaction information model included in all transactions.
type TransactionInfo struct {
	// The block height in which the transaction was included.
	Height uint64 `json:"height,omitempty"`
	// The index representing either transaction index/position within block or
	// within an aggregate transaction.
	Index int `json:"index,omitempty"`
	// The transaction db id.
	Id *coders.UInt64 `json:"id,omitempty"`
	// The transaction hash.
	Hash string `json:"hash,omitempty"`
	// The transaction merkle hash.
	MerkleComponentHash string `json:"merkleComponentHash,omitempty"`
}
