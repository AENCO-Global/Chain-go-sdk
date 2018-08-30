package main

import (
	"fmt"
	"github.com/slackve/nem2-sdk-go/sdk/model/account"
	"github.com/slackve/nem2-sdk-go/sdk/model/blockchain"

	"github.com/slackve/nem2-sdk-go/sdk/utils"
)

// Private key generated with catapult.tools.address
//
// bin/catapult.tools.address -g 1 --network mijin-test
// --- generating 1 keys ---
//           private key: 6C07F78D8C932626F6550FB114C26EFAFE2EC40220E44E1EF0180D9FB89A0AF0
//            public key: 8C2C06CCCDDFBC964345C051B3A94906813DCB198BA8A56378DA6ED1D2E99B58
//  address (mijin-test): SCIXWKDPLL7L7IQO6LWN2HJXHYUV3F7ZQX3NUKM3

const (
	private = "6C07F78D8C932626F6550FB114C26EFAFE2EC40220E44E1EF0180D9FB89A0AF0"
	public  = "8C2C06CCCDDFBC964345C051B3A94906813DCB198BA8A56378DA6ED1D2E99B58"
)

func main() {

	// The Account structure describes an account private key, public key and address.
	// Param private - A a private key in hex or an empty string.
	// Param networktype - A blockchain.NetworkType:
	// MAIN_NET		= Main net network.
	// TEST_NET		= Test net network.
	// MIJIN		= Mijin net network.
	// MIJIN_TEST	= Mijin test net network.
	// Return - A Account struct
	keypair, _ := account.CreateFromPrivateKey(private, blockchain.NetworkType.MIJIN_TEST)

	fmt.Println("-- Generating 1 keys ---")
	fmt.Printf("\t\t private key: %v\n", utils.Bt2Hex(keypair.KeyPair.PrivateKey))
	fmt.Printf("\t\t  public key: %v\n", utils.Bt2Hex(keypair.KeyPair.PublicKey))
	fmt.Printf("address (mijin-test): %v\n\n", keypair.Address.Address)

	// The Address structure describes an address with its network.
	// Create an Account from a given public key.
	// Param public - A a public key in hex or an empty string.
	// Return - A Address struct.
	account, err := account.CreateFromPublicKey(public, blockchain.NetworkType.MIJIN_TEST)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Address struct: %+v\n", account)

}
