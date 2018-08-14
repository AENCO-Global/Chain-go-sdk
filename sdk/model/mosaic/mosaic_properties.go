package mosaic

// Mosaic properties model
type MosaicProperties struct {
	// The divisibility determines up to what decimal place the mosaic can be divided into.
	// Thus a divisibility of 3 means that a mosaic can be divided into smallest parts of 0.001 mosaics
	// i.e. milli mosaics is the smallest sub-unit.
	// When transferring mosaics via a transfer transaction the quantity transferred
	// is given in multiples of those smallest parts.
	// The divisibility must be in the range of 0 and 6. The default value is "0".
	Divisibility int
	// The duration in blocks a mosaic will be available.
	// After the duration finishes mosaic is inactive and can be renewed.
	Duration uint64
	// The creator can choose between a definition that allows a mosaic supply change at a
	// later point or an immutable supply.
	// Allowed values for the property are "true" and "false". The default value is "false".
	SupplyMutable bool
	// The creator can choose if the mosaic definition should allow for transfers of the
	// mosaic among accounts other than the creator.
	// If the property 'transferable' is set to "false", only transfer transactions
	// having the creator as sender or as recipient can transfer mosaics of that type.
	// If set to "true" the mosaics can be transferred to and from arbitrary accounts.
	// Allowed values for the property are thus "true" and "false". The default value is "true".
	Transferable bool
	// Levy mutable.
	LevyMutable bool
}
