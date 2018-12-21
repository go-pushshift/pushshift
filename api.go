package pushshift

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var Client http.Client

type SubredditType uint
const (
	SubredditPublic = SubredditType(iota)
	SubredditPrivate
	SubredditGoldOnly
	SubredditEmployeesOnly
	SubredditUser
	SubredditQuarantined
)

type GildedType uint
const (
	NotGilded      = GildedType(0)
	GildedSilver   = GildedType(1)
	GildedGold     = GildedType(2)
	GildedPlatinum = GildedType(4)
	Gilded         = GildedType(7)
)

const timeStampFormat = "2006-01-02 15:04:05"

func (t SubredditType) String() string {
	switch t {
	case SubredditPublic:
		return "public"
	case SubredditPrivate:
		return "private"
	case SubredditGoldOnly:
		return "gold_only"
	case SubredditEmployeesOnly:
		return "employees_only"
	case SubredditUser:
		return "user"
	case SubredditQuarantined:
		return "quarantined"
	default:
		panic("Invalid subreddit type")
	}
}

func addQueryList(u url.Values, k string, v []string) {
	news := strings.Join(v, ",")
	prevs := u[k]
	if len(prevs) == 0 {
		u[k] = []string{ news }
	} else if len(prevs) == 1 {
		if len(prevs[0]) == 0 {
			u[k][0] = news
		} else {
			u[k][0] = prevs[0] + "," + news
		}
	}
}

func GreaterThan(i int64) int64 {
	return i | 0x2000000000000000
}

func LessThan(i int64) int64 {
	return i | 0x4000000000000000
}

func setQueryNumeric(u url.Values, key string, num int64) {
	highBits := num >> 61

	switch highBits {
	case 0, 4:
		// Identity number
		u.Set(key, fmt.Sprintf("%d", num))
	case 1, 5:
		// Greater-than number
		u.Set(key, fmt.Sprintf(">%d", num & 0x9FFFFFFFFFFFFFFF))
	case 2, 6:
		// Positive less-than number
		u.Set(key, fmt.Sprintf("<%d", num & 0x9FFFFFFFFFFFFFFF))
	default:
		panic("integer overflow")
	}
}

func setHTTPHeaders(h *http.Header) {
	h.Set("User-Agent", "go-pushshift-api")
}
