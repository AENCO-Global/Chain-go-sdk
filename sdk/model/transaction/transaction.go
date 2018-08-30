package transaction

import (
	"github.com/slackve/nem2-sdk-go/core/coders"
)

// An transaction struct that serves as the base class of all NEM transactions.
type TransactionStc struct {
	// The transaction type.
	Type int `json:"type,omitempty"`
	// The network type.
	NetworkType int `json:"networkType,omitempty"`
	// The transaction version number.
	Version int `json:"version,omitempty"`
	// The deadline to include the transaction.
	Deadline *Deadline `json:"deadline,omitempty"`
	// The fee for the transaction. The higher the fee, the higher the priority of the transaction.
	// Transactions with high priority get included in a block before transactions with lower priority.
	Fee coders.UInt64 `json:"fee,omitempty"`
	// The transaction signature (missing if part of an aggregate transaction).
	Signature string `json:"signature,omitempty"`
	// The account of the transaction creator.
	Signer *PublicAccount `json:"signer,omitempty"`
	// Transactions meta data object contains additional information about the transaction.
	TransactionInfo *TransactionInfo `json:"transactionInfo,omitempty"`
}

var Transaction = struct{
	Create func(int, int, int, Deadline) TransactionStc
}{
	Create: CreateTransaction,
}

func CreateTransaction(Type, networkType, version int, deadline Deadline) TransactionStc {
	return TransactionStc{
		Type:            Type,
		NetworkType:     networkType,
		Version:         version,
		Deadline:        &deadline,
		Fee:             coders.UInt64{},
		Signature:       "",
		Signer:          nil,
		TransactionInfo: nil,
	}
}