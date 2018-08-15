package infrastructure

type MosaicDefinitionDto struct {
	NamespaceId *UInt64Dto `json:"namespaceId"`

	MosaicId *UInt64Dto `json:"mosaicId"`

	Supply *UInt64Dto `json:"supply"`

	Height *UInt64Dto `json:"height"`

	Owner string `json:"owner"`

	Properties *[]MosaicPropertiesDto `json:"properties"`

	Levy *interface{} `json:"levy"`
}
