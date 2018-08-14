package transaction

// SignedTransaction struct is used to transfer the transaction data and the signature to NIS
// in order to initiate and broadcast a transaction.
type SignedTransaction struct {
	// Transaction serialized data
	Payload string
	// Transaction hash
	Hash string
	// Transaction signer
	Signer string
	// Transaction type
	Type int
	// Signer network type
	NetworkType int
}
