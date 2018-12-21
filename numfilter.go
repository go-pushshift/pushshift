package pushshift

import "fmt"

// 000: Positive match
// 111: Negative match
// 001: Positive gt
// 110: Negative gt
// 010: Positive lt
// 101: Negative lt

// NumFilter is a fuzzy numeric constraint
// used as a parameter in various filter functions.
//
// To match all comments with a score of exactly 42:
//   query.Score(42)
// Matching comments with a score greater or less than 42 is also easy:
//   query.Score(pushshift.GreaterThan(42))
//   query.Score(pushshift.LessThan(42))
type NumFilter int64

// Returns a new numeric filter that matches
// values greater than the specified integer.
//   filter := pushshift.GreaterThan(3)
//   filter.Matches(2) → false
func GreaterThan(n int64) NumFilter {
	n = int64(uint64(n) ^ 0x2000000000000000)
	return NumFilter(n)
}

// Returns a new numeric filter that matches
// values less than the specified integer.
//   filter := pushshift.LessThan(3)
//   filter.Matches(2) → true
func LessThan(n int64) NumFilter {
	n = int64(uint64(n) ^ 0x4000000000000000)
	return NumFilter(n)
}

// Check if an integer matches the filter.
func (n NumFilter) Matches(i int64) bool {
	highBits := uint64(n) >> 61

	switch highBits {
	case 0, 7:
		// Identity number
		return i == int64(n)
	case 1:
		// Positive greater-than number
		return i > int64(uint64(n)&0x9FFFFFFFFFFFFFFF)
	case 2:
		// Positive less-than number
		return i < int64(uint64(n)&0x9FFFFFFFFFFFFFFF)
	case 6:
		// Negative greater-than number
		return i > int64(uint64(n)|0xE000000000000000)
	case 5:
		// Negative less-than number
		return i < int64(uint64(n)|0xE000000000000000)
	default:
		// Invalid mode
		return false
	}
}

// Format the numeric filter as a PushShift-compatible string.
//   NumFilter(3).String() → "3"
//   NumFilter(3).OrLess().String() → "<4"
//   NumFilter(3).OrMore().String() → ">2"
func (n NumFilter) String() string {
	highBits := uint64(n) >> 61

	switch highBits {
	case 0, 7:
		// Identity number
		return fmt.Sprintf("%d", n)
	case 1:
		// Positive greater-than number
		return fmt.Sprintf(">%d", int64(uint64(n)&0x9FFFFFFFFFFFFFFF))
	case 2:
		// Positive less-than number
		return fmt.Sprintf("<%d", int64(uint64(n)&0x9FFFFFFFFFFFFFFF))
	case 6:
		// Negative greater-than number
		return fmt.Sprintf(">%d", int64(uint64(n)|0xE000000000000000))
	case 5:
		// Negative less-than number
		return fmt.Sprintf("<%d", int64(uint64(n)|0xE000000000000000))
	default:
		// Invalid mode
		return ""
	}
}
