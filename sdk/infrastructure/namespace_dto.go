package infrastructure

type NamespaceDto struct {
	Type_ int32 `json:"type"`

	Depth int32 `json:"depth"`

	Level0 *UInt64Dto `json:"level0"`

	Level1 *UInt64Dto `json:"level1,omitempty"`

	Level2 *UInt64Dto `json:"level2,omitempty"`

	ParentId *UInt64Dto `json:"parentId"`

	Owner string `json:"owner"`

	OwnerAddress string `json:"ownerAddress,omitempty"`

	StartHeight *UInt64Dto `json:"startHeight"`

	EndHeight *UInt64Dto `json:"endHeight"`
}
