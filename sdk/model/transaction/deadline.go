package transaction

import "time"

const TimestampNemesisBlock = 1459468800

// The deadline of the transaction.
// The deadline is given as the number of seconds elapsed since the creation of the nemesis block.
// If a transaction does not get included in a block before the deadline is reached, it is deleted.

type Deadline struct {
	Value int64 `json:"Value,omitempty"`
}

// Create deadline model
// returns Deadline struct
func DeadLine() Deadline {
	deadline := time.Now().UnixNano() / int64(time.Millisecond)
	deadline = deadline - (TimestampNemesisBlock * 1000)
	return Deadline{
		Value: deadline,
	}
}
