package transaction

// An transaction struct that serves as the base class of all NEM transactions.
type Transaction struct {
	// The transaction type.
	Type int
	// The network type.
	NetworkType int
	// The transaction version number.
	Version int
	// The deadline to include the transaction.
	Deadline Deadline
	// The fee for the transaction. The higher the fee, the higher the priority of the transaction.
	// Transactions with high priority get included in a block before transactions with lower priority.
	Fee uint64
	// The transaction signature (missing if part of an aggregate transaction).
	Signature string
	// The account of the transaction creator.
	Signer PublicAccount
	// Transactions meta data object contains additional information about the transaction.
	TransactionInfo TransactionInfo
}
