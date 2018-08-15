package transaction

// The address structure describes an address with its network.
type Address struct {
	// The address value.
	Address string
	// The NEM network type.
	NetworkType int
}
