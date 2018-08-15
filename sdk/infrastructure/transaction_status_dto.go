package infrastructure

type TransactionStatusDto struct {
	Group string `json:"group,omitempty"`

	Status string `json:"status"`

	Hash string `json:"hash,omitempty"`

	Deadline *UInt64Dto `json:"deadline,omitempty"`

	Height *UInt64Dto `json:"height,omitempty"`
}
