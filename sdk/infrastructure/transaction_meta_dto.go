package infrastructure

type TransactionMetaDto struct {
	Hash                string `json:"hash"`
	Height              []int  `json:"height"`
	ID                  string `json:"id"`
	Index               int    `json:"index"`
	MerkleComponentHash string `json:"merkleComponentHash"`
}
