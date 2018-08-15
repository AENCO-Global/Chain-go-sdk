package infrastructure

type TransactionInfoDto struct {

	Meta *TransactionMetaDto `json:"meta,omitempty"`

	Transaction *TransactionDto `json:"transaction,omitempty"`
}
