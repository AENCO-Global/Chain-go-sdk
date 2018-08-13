package blockchain

// The blockchain score structure describes blockchain difficulty.
type BlockchainScore struct {
	// Low part of the blockchain score.
	ScoreLow uint64
	// High part of the blockchain score.
	ScoreHigh uint64
}