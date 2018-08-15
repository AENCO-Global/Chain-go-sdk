package infrastructure

type MosaicNameDto struct {
	ParentId *UInt64Dto `json:"parentId"`

	MosaicId *UInt64Dto `json:"mosaicId"`

	Name string `json:"name"`
}
