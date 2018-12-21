package pushshift

import (
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
		u[k] = []string{news}
	} else if len(prevs) == 1 {
		if len(prevs[0]) == 0 {
			u[k][0] = news
		} else {
			u[k][0] = prevs[0] + "," + news
		}
	}
}

func setHTTPHeaders(h *http.Header) {
	h.Set("User-Agent", "go-pushshift-api")
}
