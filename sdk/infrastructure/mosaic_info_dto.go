package infrastructure

type MosaicInfoDto struct {

	Meta *NamespaceMosaicMetaDto `json:"meta"`

	Mosaic *MosaicDefinitionDto `json:"mosaic"`
}
