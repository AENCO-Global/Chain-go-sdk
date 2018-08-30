package transaction

import (
	"github.com/slackve/nem2-sdk-go/sdk/model/account"
	"github.com/slackve/nem2-sdk-go/sdk/model/mosaic"
)

// Transfer transactions contain data about transfers of mosaics and message to another account.
type TransferTransactionStc struct {
	// extends Transaction.
	TransactionStc
	// The address of the recipient.
	Recipient *account.Address `json:"recipient,omitempty"`
	// The array of Mosaic objects.
	Mosaics *[]mosaic.MosaicStc `json:"mosaics,omitempty"`
	// The transaction message of 2048 characters.
	Message *Message `json:"message,omitempty"`
}

var TransferTransaction = struct {
	Create func(Deadline, account.Address, []mosaic.MosaicStc, Message, int) TransferTransactionStc
}{
	Create: CreateTransferTransaction,
}

// Create a transfer transaction object
// param deadline - The deadline to include the transaction.
// param recipient - The recipient of the transaction.
// param mosaics - The array of mosaics.
// param message - The transaction message.
// param networkType - The network type.
// returns TransferTransaction struct
func CreateTransferTransaction(deadline Deadline, recipient account.Address, mosaics []mosaic.MosaicStc,
	message Message, networkType int) TransferTransactionStc {

	return TransferTransactionStc{
		TransactionStc: Transaction.Create(TransactionType.TRANSFER, networkType, 3, deadline),
		Recipient:      &recipient,
		Mosaics:        &mosaics,
		Message:        &message,
	}
}
