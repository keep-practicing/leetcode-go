package utils

const (
	// MaxUint max unsigned int.
	MaxUint = ^uint(0)
	// MaxInt max int.
	MaxInt = int(MaxUint >> 1)

	// MinUint min unsigned int.
	MinUint = 0

	// MinInt min int
	MinInt = -MaxInt - 1
)
