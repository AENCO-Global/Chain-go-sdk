package model

import "github.com/slackve/nem2-sdk-go/core/coders"

func FromUint(value int64) coders.UInt64 {
	return coders.FromUint(value)
}

func ToDto(value []int) coders.UInt64 {
	if len(value) != 2 || value[0] < 0 || value[1] < 0 {
		panic("uintArray must be be an array of two uint numbers")
	}
	return coders.UInt64{
		Lower:  uint64(value[1]),
		Higher: uint64(value[0]),
	}
}
