package infrastructure

type BlockMetaDto struct {
	Hash string `json:"hash,omitempty"`

	GenerationHash string `json:"generationHash,omitempty"`

	TotalFee *[]UInt64Dto `json:"totalFee,omitempty"`

	MerkleTree []string `json:"-"`

	NumTransactions float32 `json:"numTransactions,omitempty"`
}
