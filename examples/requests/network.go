package main

import (
	"fmt"
	"github.com/slackve/nem2-sdk-go/sdk/infrastructure"
	"github.com/slackve/nem2-sdk-go/sdk/utils"
)

func main() {

	// Network routes.
	const Server = "http://catapult.isarq.com:3000"

	client := infrastructure.NewAPIClient()
	client.ChangeServer(Server)

	// 01 - Get the current network type of the chain.
	a, err := client.NetworkRoutesApi.GetNetworkType()
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetNetworkType:\n%s", utils.Struc2Json(a))
}
