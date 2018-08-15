package infrastructure

type BlockDto struct {

	Signature string `json:"signature,omitempty"`

	Signer string `json:"signer,omitempty"`

	Version float32 `json:"version,omitempty"`

	Type_ float32 `json:"type"`

	Height *[]UInt64Dto `json:"height,omitempty"`

	Timestamp *[]UInt64Dto `json:"timestamp,omitempty"`

	Difficulty *[]UInt64Dto `json:"difficulty,omitempty"`

	PreviousBlockHash string `json:"previousBlockHash,omitempty"`

	BlockTransactionsHash string `json:"blockTransactionsHash,omitempty"`
}
