package infrastructure

type NamespaceInfoDto struct {

	Meta *NamespaceMosaicMetaDto `json:"meta"`

	Namespace *NamespaceDto `json:"namespace"`
}
