package coders

import (
	"encoding/base32"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/ripemd160"
	. "github.com/slackve/nem2-sdk-go/sdk/model/blockchain"
	"github.com/slackve/nem2-sdk-go/sdk/utils"
)

var Sizes = struct {
	Ripemd160      int
	AddressDecoded int
	AddressEncoded int
	Key            int
	Checksum       int
}{
	Ripemd160:      20,
	AddressDecoded: 25,
	AddressEncoded: 40,
	Key:            32,
	Checksum:       4,
}

type AddressRaw []byte
// The address structure describes an address with its network.
// param address - A Address in plain format
// param networkType - A Network type
// return error
func Address(address string, networkType int) error {
	address = utils.AddressFormat(address)
	addressNetwork := address[0]
	if networkType == NetworkType.MAIN_NET && addressNetwork != 'N' {
		return errors.New("MAIN_NET Address must start with N")
	} else if networkType == NetworkType.TEST_NET && addressNetwork != 'T' {
		return errors.New("TEST_NET Address must start with T")
	} else if networkType == NetworkType.MIJIN && addressNetwork != 'M' {
		return errors.New("MIJIN Address must start with M")
	} else if networkType == NetworkType.MIJIN_TEST && addressNetwork != 'S' {
		return errors.New("MIJIN_TEST Address must start with S")
	}
	return nil
}

// Converts a public key to a decoded address for a specific network.
// param Key - A Publickey
// param networkType - A networkIdentifier
// return The decoded address corresponding to the inputs.
func PublicKeyToAddress(Key string, networkType int) AddressRaw {
	// step 1: sha3 hash of the public key
	publicKeyHash := sha3.Sum256(utils.Hex2Bt(Key))
	//fmt.Println("publicKeyHash:", publicKeyHash)

	// step 2: ripemd160 hash of (1)
	r := ripemd160.New()
	r.Write(publicKeyHash[:])
	ripemdHash := r.Sum(nil)

	//// step 3: add network identifier byte in front of (2)
	decodedAddress := make([]byte, Sizes.AddressDecoded)
	decodedAddress[0] = byte(networkType)
	copy(decodedAddress[1:], ripemdHash)

	// step 4: get the checksum of (3)
	stepThreeChecksum := generateChecksum(decodedAddress[:21])

	// step 5: concatenate (3) and (4)
	copy(decodedAddress[len(decodedAddress)-4:], stepThreeChecksum)

	// step 6: base32 encode (5)
	return decodedAddress
}

// Converts an encoded address string to a decoded address.
// param encoded The encoded address string.
// returns The decoded address corresponding to the input.
func StringToAddress(encoded string) ([]byte, error) {
	if Sizes.AddressEncoded != len(encoded) {
		panic(fmt.Errorf("%v does not represent a valid encoded address", encoded))
	}
	return base32.StdEncoding.DecodeString(encoded)
}

// Converts a decoded address to an encoded address string.
// param decoded The decoded address.
// returns The encoded address string corresponding to the input.
func AddressToString(decoded []byte) string {
	if Sizes.AddressDecoded != len(decoded) {
		panic(fmt.Errorf("%v does not represent a valid decoded address", utils.Bt2Hex(decoded)))
	}
	return base32.StdEncoding.EncodeToString(decoded)
}

// Determines the validity of a decoded address.
// param decoded The decoded address.
// returns true if the decoded address is valid, false otherwise.
func IsValidAddress() {
	hash := sha3.New256()
	//checksumBegin := Sizes.AddressDecoded - Sizes.Checksum
	hash.Sum(nil)

}

// Determines the validity of an encoded address string.
// param encoded The encoded address string.
// returns true if the encoded address string is valid, false otherwise.
//func IsValidEncodedAddress(encoded string) bool {
//	if Sizes.AddressEncoded != len(encoded){
//		return false
//	}
//	decoded, err := StringToAddress(encoded)
//	if err != nil{
//		return false
//	}
//	IsValidAddress(decoded)
//	return true
//}

func generateChecksum(input []byte) []byte {
	// step 1: sha3 hash of (input)
	sha3StepThreeHash := sha3.Sum256(input)

	// step 2: get the first X bytes of (1)
	return sha3StepThreeHash[0:Sizes.Checksum]
}
