package mosaic

import (
	"github.com/slackve/nem2-sdk-go/core/coders"
	"github.com/slackve/nem2-sdk-go/sdk/model"
)

var Mosaic = struct {
	New func(string, int64) MosaicStc
}{
	New: CreateMosaic,
}
// A mosaic describes an instance of a mosaic definition.
// Mosaics can be transferred by means of a transfer transaction.
type MosaicStc struct {
	// The mosaic id
	Id MosaicId `json:"id,omitempty"`
	// The mosaic amount. The quantity is always given in smallest units for the mosaic
	// i.e. if it has a divisibility of 3 the quantity is given in millis.
	Amount coders.UInt64 `json:"amount,omitempty"`
}

func CreateMosaic(mosaicId string, amount int64) MosaicStc {
	return MosaicStc{
		Id:     MosaicId{
			Id:       coders.UInt64{
				Lower:  0,
				Higher: 0,
			},
			FullName: mosaicId,
		},
		Amount: model.FromUint(amount),
	}
}