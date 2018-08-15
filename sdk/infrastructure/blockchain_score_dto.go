package infrastructure

type BlockchainScoreDto struct {

	ScoreHigh *[]UInt64Dto `json:"scoreHigh"`

	ScoreLow *[]UInt64Dto `json:"scoreLow"`
}
