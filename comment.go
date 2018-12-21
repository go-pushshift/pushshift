package pushshift

import (
	"net/url"
	"strconv"
	"time"
)

const commentSearchBaseURL = "https://beta.pushshift.io/reddit/comment/search/"

type CommentSearch struct {
	v url.Values
}

func SearchComments(q string) *CommentSearch {
	c := new(CommentSearch)
	c.v.Set("q", q)
	return c
}

func (q *CommentSearch) Subreddit(subs ...string) *CommentSearch {
	addQueryList(q.v, "subreddit", subs)
	return q
}

func (q *CommentSearch) WithoutSubreddit(subs ...string) *CommentSearch {
	for i, sub := range subs {
		subs[i] = "!" + sub
	}
	addQueryList(q.v, "subreddit", subs)
	return q
}

func (q *CommentSearch) Author(authors ...string) *CommentSearch {
	addQueryList(q.v, "author", authors)
	return q
}

func (q *CommentSearch) WithoutAuthor(authors ...string) *CommentSearch {
	for i, author := range authors {
		authors[i] = "!" + author
	}
	addQueryList(q.v, "author", authors)
	return q
}

func (q *CommentSearch) Before(t time.Time) *CommentSearch {
	q.v.Set("before", t.UTC().Format(timeStampFormat))
	return q
}

func (q *CommentSearch) After(t time.Time) *CommentSearch {
	q.v.Set("after", t.UTC().Format(timeStampFormat))
	return q
}

func (q *CommentSearch) LinkID(links ...string) *CommentSearch {
	addQueryList(q.v, "link_id", links)
	return q
}

func (q *CommentSearch) Sticky(sticky bool) *CommentSearch {
	q.v.Set("stickied", strconv.FormatBool(sticky))
	return q
}

func (q *CommentSearch) CanGild(canGild bool) *CommentSearch {
	q.v.Set("can_gild", strconv.FormatBool(canGild))
	return q
}

func (q *CommentSearch) ByBot(isBot bool) *CommentSearch {
	q.v.Set("is_bot", strconv.FormatBool(isBot))
	return q
}

func (q *CommentSearch) NestLevel(i NumFilter) *CommentSearch {
	q.v.Set("nest_level", i.String())
	return q
}

func (q *CommentSearch) Edited(edited bool) *CommentSearch {
	q.v.Set("is_edited", strconv.FormatBool(edited))
	return q
}

func (q *CommentSearch) RemovedByModerator(removed bool) *CommentSearch {
	q.v.Set("mod_removed", strconv.FormatBool(removed))
	return q
}

func (q *CommentSearch) RemovedByUser(removed bool) *CommentSearch {
	q.v.Set("user_removed", strconv.FormatBool(removed))
	return q
}

func (q *CommentSearch) ReplyDelay(seconds NumFilter) *CommentSearch {
	q.v.Set("reply_delay", seconds.String())
	return q
}

func (q *CommentSearch) Length(i NumFilter) *CommentSearch {
	q.v.Set("length", i.String())
	return q
}

func (q *CommentSearch) Score(i NumFilter) *CommentSearch {
	q.v.Set("score", i.String())
	return q
}

func (q *CommentSearch) Controversial(controv bool) *CommentSearch {
	q.v.Set("controversiality", strconv.FormatBool(controv))
	return q
}

func (q *CommentSearch) IsNoFollow(noFollow bool) *CommentSearch {
	q.v.Set("no_follow", strconv.FormatBool(noFollow))
	return q
}

func (q *CommentSearch) WantsNotifications(send bool) *CommentSearch {
	q.v.Set("send_replies", strconv.FormatBool(send))
	return q
}

func (q *CommentSearch) ParentID(id string) *CommentSearch {
	q.v.Set("parent_id", id)
	return q
}

func (q *CommentSearch) ID(ids ...string) *CommentSearch {
	addQueryList(q.v, "id", ids)
	return q
}

func (q *CommentSearch) Distinguished(dist bool) *CommentSearch {
	q.v.Set("distinguished", strconv.FormatBool(dist))
	return q
}

func (q *CommentSearch) SubredditType(typ SubredditType) *CommentSearch {
	q.v.Set("subreddit_type", typ.String())
	return q
}

func (q *CommentSearch) GildedSilverTimes(i NumFilter) *CommentSearch {
	q.v.Set("gid_1", i.String())
	return q
}

func (q *CommentSearch) GildedGoldTimes(i NumFilter) *CommentSearch {
	q.v.Set("gid_2", i.String())
	return q
}

func (q *CommentSearch) GildedPlatinumTimes(i NumFilter) *CommentSearch {
	q.v.Set("gid_3", i.String())
	return q
}

func (q *CommentSearch) GildedTimes(i NumFilter) *CommentSearch {
	q.v.Set("gilded", i.String())
	return q
}

func (q *CommentSearch) IsGilded(gilded bool) *CommentSearch {
	if gilded {
		return q.GildedTimes(GreaterThan(0))
	} else {
		return q.GildedTimes(0)
	}
}

func (q *CommentSearch) AuthorFlair(flair string) *CommentSearch {
	q.v.Set("author_flair_text", flair)
	return q
}

func (q *CommentSearch) LinkFlair(flair string) *CommentSearch {
	q.v.Set("link_flair_text", flair)
	return q
}

func (q *CommentSearch) AsChart() *ChartQuery {
	q.v.Set("output", "png")
	return &ChartQuery{
		v:       q.v,
		baseURL: commentSearchBaseURL,
	}
}
