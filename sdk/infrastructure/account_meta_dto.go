package infrastructure

type AccountMetaDto struct {
	Meta *AccountMetaDto `json:"meta"`

	Account *AccountDto `json:"account"`
}
