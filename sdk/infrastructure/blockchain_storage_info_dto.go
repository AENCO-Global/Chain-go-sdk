package infrastructure

type BlockchainStorageInfoDto struct {
	NumBlocks int32 `json:"numBlocks"`

	NumTransactions int32 `json:"numTransactions"`

	NumAccounts int32 `json:"numAccounts"`
}
