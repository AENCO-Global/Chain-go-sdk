package transaction

import "time"

// The deadline of the transaction.
// The deadline is given as the number of seconds elapsed since the creation of the nemesis block.
// If a transaction does not get included in a block before the deadline is reached, it is deleted.

type Deadline struct {
	Value uint64
}
// Create deadline model
// param deadline
// param chronoUnit
// returns Deadline struct
func (d *Deadline) Create() Deadline {
	deadline := 2
	chronoUnit := time.Now().Hour()
	networkTimeStamp := time.Now()
}
