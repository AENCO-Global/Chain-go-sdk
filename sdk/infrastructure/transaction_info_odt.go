package infrastructure

type TransactionInfoDto struct {

	Meta *TransactionMetaDto `json:"meta"`

	Transaction *TransactionDto `json:"transaction"`
}
