package coders


type UInt64 struct {
	// uint64 lower part
	Lower uint64 `json:"lower"`
	// uint64 higher part
	Higher uint64 `json:"higher"`
}

// Converts a numeric unsigned integer into a uint64.
// param number The unsigned integer.
// returns {module:coders/uint64~uint64} The uint64 representation of the input.
func FromUint(number int64) UInt64 {
	value := uint64(number)
	return UInt64{(value & 0xFFFFFFFF) >> 0, (value / 0x100000000) >> 0}
}