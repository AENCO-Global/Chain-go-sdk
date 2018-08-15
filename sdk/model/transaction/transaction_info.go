package transaction

// Transaction information model included in all transactions.
type TransactionInfo struct {
	// The block height in which the transaction was included.
	Height uint64 `json:"height"`
	// The index representing either transaction index/position within block or
	// within an aggregate transaction.
	Index int `json:"index"`
	// The transaction db id.
	Id string `json:"id"`
	// The transaction hash.
	Hash string `json:"hash"`
	// The transaction merkle hash.
	MerkleComponentHash string `json:"merkleComponentHash"`
}
