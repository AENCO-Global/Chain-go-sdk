package main

import (
	"fmt"
	"github.com/slackve/nem2-sdk-go/sdk/infrastructure"
	"github.com/slackve/nem2-sdk-go/sdk/utils"
)

const (
	TxID = "5B2B0F61415CD864572BDA46"
	Hash = "306B7F09B6F62AC86F9CC32B07DF129CE2269FB2A669CC705776F5A553848582"
)

func main() {

	// Transactions routes.
	const Server = "http://catapult.isarq.com:3000"

	client := infrastructure.NewAPIClient()
	client.ChangeServer(Server)

	// 01 - Get transaction information.
	// Param Hash - transactionId or hash.
	a, err := client.TransactionRoutesApi.GetTransaction(Hash)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetTransaction:\n%s", utils.Struc2Json(a))

	// 02 - Get transaction information for a given set of transactionId or hash.
	// Param Hash - transactionId or hash.
	var TxIds infrastructure.TransactionIds
	// Append attachment into TransactionId
	TxIds.TransactionIds = append(TxIds.TransactionIds, TxID)
	// Append attachment into TransactionId
	//TxIds.TransactionIds = append(TxIds.TransactionIds, TxID2)

	b, err := client.TransactionRoutesApi.GetTransactions(TxIds)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetTransactions:\n%s", utils.Struc2Json(b))

	// 03 - Get transaction status.
	// Param Hash - Transaction hash.
	c, err := client.TransactionRoutesApi.GetTransactionStatus(Hash)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetTransactionStatus:\n%s", utils.Struc2Json(c))

	// 04 - Get an array of transaction statuses for a given set of transactionId or hash.
	// Param Hash - Slice of transactionIds or hashes.
	var TxHash infrastructure.TransactionHashes
	// Append attachment into TransactionHash
	TxHash.Hashes = append(TxHash.Hashes, Hash)
	// Append attachment into TransactionHash
	//TxIds.Hashes = append(TxIds.Hashes, Hash2)

	d, err := client.TransactionRoutesApi.GetTransactionsStatuses(TxHash)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetTransactionsStatuses:\n%s", utils.Struc2Json(d))
}
