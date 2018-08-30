package transaction

import "github.com/slackve/nem2-sdk-go/sdk/model/account"

// The address structure describes an address with its network.
type Address struct {
	account.Address
}
