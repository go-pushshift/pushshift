package pushshift

import (
	"sync"
	"time"
)

const queryWaitTime = time.Second

var rateLimitConf struct {
	sync.Mutex
	lastQuery time.Time
}

func rateLimit() {
	rateLimitConf.Lock()
	defer rateLimitConf.Unlock()

	now := time.Now()
	sinceLastQuery := time.Since(rateLimitConf.lastQuery)
	time.Sleep(queryWaitTime - sinceLastQuery)

	rateLimitConf.lastQuery = now
}
