package infrastructure

type AccountInfoDto struct {

	Meta *AccountMetaDto `json:"meta"`

	Account *AccountDto `json:"account"`
}
