package infrastructure

type BlockInfoDto struct {

	Meta *BlockMetaDto `json:"meta,omitempty"`

	Block *BlockDto `json:"block,omitempty"`
}
