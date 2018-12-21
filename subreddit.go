package pushshift

import (
	"net/url"
	"time"
)

type SubredditSearch struct {
	v url.Values
}

func SearchSubreddits(q string) *SubredditSearch {
	c := new(SubredditSearch)
	c.v.Set("q", q)
	return c
}

func (q *SubredditSearch) Subreddit(subs ...string) *SubredditSearch {
	addQueryList(q.v, "subreddit", subs)
	return q
}

func (q *SubredditSearch) WithoutSubreddit(subs ...string) *SubredditSearch {
	for i, sub := range subs {
		subs[i] = "!" + sub
	}
	addQueryList(q.v, "subreddit", subs)
	return q
}

func (q *SubredditSearch) Before(t time.Time) *SubredditSearch {
	q.v.Set("before", t.UTC().Format(timeStampFormat))
	return q
}

func (q *SubredditSearch) After(t time.Time) *SubredditSearch {
	q.v.Set("after", t.UTC().Format(timeStampFormat))
	return q
}

func (q *SubredditSearch) ID(ids ...string) *SubredditSearch {
	addQueryList(q.v, "id", ids)
	return q
}
