package main

import (
	"fmt"
	"github.com/slackve/nem2-sdk-go/sdk/infrastructure"
	"github.com/slackve/nem2-sdk-go/sdk/utils"
)

const (
	namespace  = "84b3552d375ffa4b"
	namespace2 = "d525ad41d95fcf28"
	Address    = "SCFWMP2M2HP43KJYGOBDVQ3SKX3Q6HFH6HZZ6DNR"
	Address2   = "SAUUKSFBYHI57KTEXQNJWHOKXITF7T4BXON3GVTJ"
)

func main() {

	// Blockchain routes.
	const Server = "http://catapult.isarq.com:3000"

	client := infrastructure.NewAPIClient()
	client.ChangeServer(Server)

	// 01 - Get namespace information.
	// Param mosaic - Mosaic identifier.
	a, err := client.NamespaceRoutesApi.GetNamespace(namespace)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetNamespace:\n%s", utils.Struc2Json(a))

	// 02 - Get namespaces an account owns.
	// Param Address - Account address or publicKey.
	// Param 0 - The number of namespaces to return.
	// Param "" - Identifier of the namespace after which we want the transactions to be returned.
	b, err := client.NamespaceRoutesApi.GetNamespacesFromAccount(Address, 0, "")
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetNamespacesFromAccount:\n%s", utils.Struc2Json(b))

	// 03 - Get an array of NamespaceInfo for a given set of addresses.
	// Param adds - Accounts address array.
	// Param 0 - The number of namespaces to return.
	// Param "" - Identifier of the namespace after which we want the transactions to be returned.
	var adds infrastructure.Addresses
	// Append attachment into Address
	adds.Addresses = append(adds.Addresses, Address)
	// Append attachment into Address
	adds.Addresses = append(adds.Addresses, Address2)
	c, err := client.NamespaceRoutesApi.GetNamespacesFromAccounts(adds, 0, "")
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetNamespacesFromAccounts:\n%s", utils.Struc2Json(c))

	// 04 - Get readable names for a set of namespaces.
	// Param nsp - Slice of mosaicIds.
	var nsp infrastructure.NamespaceIds
	// Append attachment into Namespace
	nsp.NamespaceIds = append(nsp.NamespaceIds, namespace)
	d, err := client.NamespaceRoutesApi.GetNamespacesNames(nsp)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetNamespacesNames:\n%s", utils.Struc2Json(d))
}
