package mosaic

// A mosaic describes an instance of a mosaic definition.
// Mosaics can be transferred by means of a transfer transaction.
type Mosaic struct {
	// The mosaic id
	Id MosaicId
	// The mosaic amount. The quantity is always given in smallest units for the mosaic
	// i.e. if it has a divisibility of 3 the quantity is given in millis.
	Amount uint64
}
