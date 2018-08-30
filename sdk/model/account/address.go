package account

import (
	"fmt"
	"github.com/slackve/nem2-sdk-go/sdk/model/blockchain"
	"github.com/slackve/nem2-sdk-go/sdk/utils"
	"github.com/slackve/nem2-sdk-go/core/coders"
)
//
// The address structure describes an address with its network.
type Address struct {
	// The address value.
	Address string `json:"address,omitempty"`
	// The NEM network type.
	NetworkType int `json:"networkType,omitempty"`
}

func CreateFromRawAddress(rawAddress string) (adress Address, err error) {
	var networkType int
	addressTrimAndUpperCase := utils.AddressFormat(rawAddress)
	if len(addressTrimAndUpperCase) != 40 {
		err = fmt.Errorf("Address %v has to be 40 characters long \n", addressTrimAndUpperCase)
	}
	switch addressTrimAndUpperCase[0] {
	case 'S':
		networkType = blockchain.NetworkType.MIJIN_TEST
	case 'M':
		networkType = blockchain.NetworkType.MIJIN
	case 'T':
		networkType = blockchain.NetworkType.TEST_NET
	case 'N':
		networkType = blockchain.NetworkType.MAIN_NET
	default:
		err = fmt.Errorf("address Network unsupported'")
	}
	return Address{
		Address:     addressTrimAndUpperCase,
		NetworkType: networkType,
	}, nil
}

func CreateFromPublicKey(publickey string, networkType int) (Address, error) {
	if !utils.IsHexadecimal(publickey){
		err := fmt.Errorf("public key must be hexadecimal only !")
		return Address{}, err
	}
	address := coders.PublicKeyToAddress(publickey, networkType)
	return  CreateFromRawAddress(coders.AddressToString(address))
}