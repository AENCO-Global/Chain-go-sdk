package account

// The public account structure contains account's address and public key.
type PublicAccount struct {
	//The account public private.
	PublicKey string
	//The account address.
	Address Address
}
