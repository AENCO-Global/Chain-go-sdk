package main

import (
	"fmt"
	"github.com/slackve/nem2-sdk-go/sdk/infrastructure"
	"github.com/slackve/nem2-sdk-go/sdk/utils"
)

func main() {

	// Blockchain routes.
	const Server = "http://catapult.isarq.com:3000"

	client := infrastructure.NewAPIClient()
	client.ChangeServer(Server)

	// 01 - Get BlockInfo for a given block height.
	// Param 1 - Block height
	a, err := client.BlockchainRoutesApi.GetBlockByHeight(1)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetBlockByHeight:\n%s", utils.Struc2Json(a))

	// 02 - Get transactions from a block.
	// Param 1 - Block height
	// Param 50 - The number of transactions to return.
	// Should be between 10 and 100, otherwise 10.
	// Param "" - Identifier of the transaction after which
	// we want the transactions to be returned Eje: "5B2B0F61415CD864572BDA5B".
	b, err := client.BlockchainRoutesApi.GetBlockTransactions(1, 50, "")
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetBlockTransactions:\n%s", utils.Struc2Json(b))

	// 03 - Get the current height of the chain.
	c, err := client.BlockchainRoutesApi.GetBlockchainHeight()
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetBlockchainHeight:\n%s", utils.Struc2Json(c))

	// 04 - Get the current score of the chain.
	d, err := client.BlockchainRoutesApi.GetBlockchainScore()
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetBlockchainScore:\n%s", utils.Struc2Json(d))

	// 05 - Get an array of BlockInfo for a given block height and limit.
	// Param 1 - Block height
	// Param 25 - Number of following blocks to be returned.
	e, err := client.BlockchainRoutesApi.GetBlocksByHeightWithLimit(1, 25)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetBlocksByHeightWithLimit:\n%s", utils.Struc2Json(e))

	// 06 - Get the storage information.
	f, err := client.BlockchainRoutesApi.GetDiagnosticStorage()
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetDiagnosticStorage:\n%s", utils.Struc2Json(f))
}
