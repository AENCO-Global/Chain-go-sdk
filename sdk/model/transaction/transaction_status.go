package transaction

// Transaction status contains basic of a transaction announced to the blockchain.
type TransactionStatus struct {
	// The transaction status group "failed", "unconfirmed", "confirmed", etc...
	Group string `json:"group"`
	// The transaction status being the error name in case of failure and success otherwise.
	Status string `json:"status"`
	// The transaction hash.
	Hash string `json:"hash"`
	// The transaction deadline.
	Deadline Deadline `json:"deadline"`
	// The height of the block at which it was confirmed or rejected.
	Height uint64 `json:"height"`
}
