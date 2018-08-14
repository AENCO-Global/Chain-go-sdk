package mosaic

import (
	"github.com/slackve/nem2-sdk-go/sdk/model/namespace"
	"github.com/slackve/nem2-sdk-go/sdk/model/account"
)

// The mosaic info structure describes a mosaic.
type MosaicInfo struct {
	// Mosaic is active.
	Active bool
	// The mosaic index.
	Index int
	// The meta data id.
	MetaId string
	// The namespace id.
	NamespaceId namespace.NamespaceId
	// The mosaic id.
	MosaicId MosaicId
	// The mosaic supply.
	Supply uint64
	// The block height were mosaic was created.
	Height uint64
	// The public key of the mosaic creator.
	Owner account.PublicAccount
	// The mosaic properties.
	Properties MosaicProperties
	// The optional levy for the mosaic. A creator can demand that
	// each mosaic transfer induces an additional fee.
	Levy interface{}
	// Mosaic divisibility.
	Divisibility int
	// Mosaic duration.
	Duration uint64
}
