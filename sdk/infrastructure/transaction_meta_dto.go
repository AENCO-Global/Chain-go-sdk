package infrastructure

type TransactionMetaDto struct {
	Hash                string `json:"hash,omitempty"`
	Height              []int  `json:"height,omitempty"`
	ID                  string `json:"id,omitempty"`
	Index               int    `json:"index,omitempty"`
	MerkleComponentHash string `json:"merkleComponentHash,omitempty"`
}
