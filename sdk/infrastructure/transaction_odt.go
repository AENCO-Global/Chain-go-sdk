package infrastructure

type TransactionDto struct {
		Deadline []int `json:"deadline"`
		Fee      *[]UInt64Dto `json:"fee"`
		Mosaics  []MosaicDto `json:"mosaics"`
		Recipient string `json:"recipient"`
		Signature string `json:"signature"`
		Signer    string `json:"signer"`
		Type      int    `json:"type"`
		Version   int    `json:"version"`
}