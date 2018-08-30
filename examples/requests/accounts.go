// Blockchain routes.

package main

import (
	"fmt"
	"github.com/slackve/nem2-sdk-go/sdk/infrastructure"
	"github.com/slackve/nem2-sdk-go/sdk/utils"
)

const (
	Address   = "SCQXJT7U2PYWP5NZUC5LSMBXCIR4WTZE5OIFJD4B"
	Address2  = "SDZAIHDGCHWWZPLHDLNTZFSKPXXGQJVBV52STHZ2"
	PublicKey = "814DA2775296E71D461CACA6CB8BF3D0485C522AAAD99C1681A9AE12C402188E"
)

func main() {

	// Account routes.
	const Server = "http://api.beta.catapult.mijin.io:3000"

	client := infrastructure.NewAPIClient()
	client.ChangeServer(Server)

	// 01 - Get AccountInfo for an account.
	// Param Address - Account address or publicKey.
	a, err := client.AccountRoutesApi.GetAccountInfo(Address)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetAccountInfo:\n%s", utils.Struc2Json(a))

	// 02 - Get AccountsInfo for different accounts.
	// Param adds - Slice of addresses..
	var adds infrastructure.Addresses
	// Append attachment into Address
	adds.Addresses = append(adds.Addresses, Address)
	// Append attachment into Address
	adds.Addresses = append(adds.Addresses, Address2)

	b, err := client.AccountRoutesApi.GetAccountsInfo(adds)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetAccountInfoBatch:\n%s", utils.Struc2Json(b))

	// 03 - Get confirmed transactions information.
	// Param PublicKey - Account publicKey.
	// Param 0 - The number of transactions to return. Should be between 10 and 100, otherwise 10.
	// Param "" - Identifier of the transaction after which we want the transactions to be returned. Eje: "5B2B0F61415CD864572BDA46".
	c, err := client.AccountRoutesApi.Transactions(PublicKey, 0, "")
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("Transactions:\n%s", utils.Struc2Json(c))

	// 04 - Get incoming transactions information.
	// Param PublicKey - Account publicKey.
	// Param 0 - The number of transactions to return. Should be between 10 and 100, otherwise 10.
	// Param "" - Identifier of the transaction after which we want the transactions to be returned. Eje: "5B2B0F61415CD864572BDA46".
	d, err := client.AccountRoutesApi.IncomingTransactions(PublicKey, 0, "")
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("IncomingTransactions:\n%s", utils.Struc2Json(d))

	// 05 - Get outgoing transactions information.
	// Param PublicKey - Account publicKey.
	// Param 0 - The number of transactions to return. Should be between 10 and 100, otherwise 10.
	// Param "" - Identifier of the transaction after which we want the transactions to be returned. Eje: "5B2B0F61415CD864572BDA46".
	e, err := client.AccountRoutesApi.OutgoingTransactions(PublicKey, 0, "")
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("OutgoingTransactions:\n%s", utils.Struc2Json(e))

	// 06 - Get unconfirmed transactions information.
	// Param PublicKey - Account publicKey.
	// Param 0 - The number of transactions to return. Should be between 10 and 100, otherwise 10.
	// Param "" - Identifier of the transaction after which we want the transactions to be returned. Eje: "5B2B0F61415CD864572BDA46".
	f, err := client.AccountRoutesApi.UnconfirmedTransactions(PublicKey, 0, "")
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("UnconfirmedTransactions:\n%s", utils.Struc2Json(f))

	// 07 - Get aggregate bonded transactions information.
	// Param PublicKey - Account publicKey.
	// Param 0 - The number of transactions to return. Should be between 10 and 100, otherwise 10.
	// Param "" - Identifier of the transaction after which we want the transactions to be returned. Eje: "5B2B0F61415CD864572BDA46".
	g, err := client.AccountRoutesApi.PartialTransactions(PublicKey, 0, "")
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("PartialTransactions:\n%s", utils.Struc2Json(g))

	// 08 - Get multisig account information.
	// Param Address - Account address or publicKey.
	h, err := client.AccountRoutesApi.GetAccountMultisig(PublicKey)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetAccountMultisig:\n%s", utils.Struc2Json(h))

	// 09 - Get multisig account graph information.
	// Param Address - Account address or publicKey.
	i, err := client.AccountRoutesApi.GetAccountMultisigGraph(PublicKey)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetAccountMultisigGraph:\n%s", utils.Struc2Json(i))
}
