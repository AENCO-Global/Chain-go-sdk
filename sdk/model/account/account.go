package account

import (
	"github.com/slackve/nem2-sdk-go/core/crypto"
	"fmt"
	"github.com/slackve/nem2-sdk-go/core/coders"
	"encoding/hex"
)

// The account structure describes an account private key, public key, address and
// allows signing transactions.
type Account struct {
	// The account address.
	Address Address
	// The account keyPair, public and private key.
	KeyPair KeyPair
}

// A catapult key pair struct composed of a public and private key.
type KeyPair struct {
	// The private key.
	PrivateKey []byte
	// The public key.
	PublicKey  []byte
}

// Create an Account from a given private key
// param privateKey - Private key from an account
// param networkType - Network type
// return - A Account struct.
func CreateFromPrivateKey(privateKey string, networkType int) (Account, error) {
	keyPair, err := crypto.KeyPairFromPrivateKey(privateKey)

	if err != nil{
		fmt.Println("ERRROR: ", err)
	}
	address := coders.PublicKeyToAddress(hex.EncodeToString(keyPair.Public), networkType)
	addr, err := CreateFromRawAddress(coders.AddressToString(address))
	return Account{
		Address: addr,
		KeyPair: KeyPair{
			PrivateKey: keyPair.Private,
			PublicKey:  keyPair.Public,
		},
	}, nil
}
