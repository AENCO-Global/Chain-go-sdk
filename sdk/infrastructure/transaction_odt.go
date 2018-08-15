package infrastructure

type TransactionDto struct {
	Deadline  *[]UInt64Dto `json:"deadline,omitempty"`
	Fee       *[]UInt64Dto `json:"fee,omitempty"`
	Mosaics   []MosaicDto  `json:"mosaics,omitempty"`
	Recipient string       `json:"recipient,omitempty"`
	Signature string       `json:"signature,omitempty"`
	Signer    string       `json:"signer,omitempty"`
	Type      int          `json:"type,omitempty"`
	Version   int          `json:"version,omitempty"`
}
