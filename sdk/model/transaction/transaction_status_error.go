package transaction

// Transaction status error model returned by listeners
type TransactionStatusError struct {
	// The transaction hash.
	Hash string `json:"hash"`
	// The status error message.
	Status string `json:"status"`
	// The transaction deadline.
	Deadline Deadline `json:"deadline"`
}
