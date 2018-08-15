package infrastructure

type NamespaceNameDto struct {
	ParentId *UInt64Dto `json:"parentId,omitempty"`

	NamespaceId *UInt64Dto `json:"namespaceId"`

	Name string `json:"name"`
}
