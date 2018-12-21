package pushshift

import (
	"net/url"
	"strconv"
	"time"
)

const submissionSearchBaseURL = "https://beta.pushshift.io/reddit/submission/search/"

type SubmissionSearch struct {
	v url.Values
}

func SearchSubmissions(q string) *SubmissionSearch {
	c := new(SubmissionSearch)
	c.v.Set("q", q)
	return c
}

func (q *SubmissionSearch) Subreddit(subs ...string) *SubmissionSearch {
	addQueryList(q.v, "subreddit", subs)
	return q
}

func (q *SubmissionSearch) WithoutSubreddit(subs ...string) *SubmissionSearch {
	for i, sub := range subs {
		subs[i] = "!" + sub
	}
	addQueryList(q.v, "subreddit", subs)
	return q
}

func (q *SubmissionSearch) Author(authors ...string) *SubmissionSearch {
	addQueryList(q.v, "author", authors)
	return q
}

func (q *SubmissionSearch) WithoutAuthor(authors ...string) *SubmissionSearch {
	for i, author := range authors {
		authors[i] = "!" + author
	}
	addQueryList(q.v, "author", authors)
	return q
}

func (q *SubmissionSearch) Domain(domains ...string) *SubmissionSearch {
	addQueryList(q.v, "domain", domains)
	return q
}

func (q *SubmissionSearch) WithoutDomain(domains ...string) *SubmissionSearch {
	for i, domain := range domains {
		domains[i] = "!" + domain
	}
	addQueryList(q.v, "domain", domains)
	return q
}

func (q *SubmissionSearch) Title(title string) *SubmissionSearch {
	q.v.Set("title", title)
	return q
}

func (q *SubmissionSearch) SelfText(text string) *SubmissionSearch {
	q.v.Set("selftext", text)
	return q
}

func (q *SubmissionSearch) Before(t time.Time) *SubmissionSearch {
	q.v.Set("before", t.UTC().Format(timeStampFormat))
	return q
}

func (q *SubmissionSearch) After(t time.Time) *SubmissionSearch {
	q.v.Set("after", t.UTC().Format(timeStampFormat))
	return q
}

func (q *SubmissionSearch) Self(self bool) *SubmissionSearch {
	q.v.Set("is_self", strconv.FormatBool(self))
	return q
}

func (q *SubmissionSearch) Locked(locked bool) *SubmissionSearch {
	q.v.Set("locked", strconv.FormatBool(locked))
	return q
}

func (q *SubmissionSearch) Spoiler(spoiler bool) *SubmissionSearch {
	q.v.Set("spoiler", strconv.FormatBool(spoiler))
	return q
}

func (q *SubmissionSearch) Over18(over18 bool) *SubmissionSearch {
	q.v.Set("over18", strconv.FormatBool(over18))
	return q
}

func (q *SubmissionSearch) BrandSafe(brandSafe bool) *SubmissionSearch {
	q.v.Set("brand_safe", strconv.FormatBool(brandSafe))
	return q
}

func (q *SubmissionSearch) ContestMode(contestMode bool) *SubmissionSearch {
	q.v.Set("contest_mode", strconv.FormatBool(contestMode))
	return q
}
func (q *SubmissionSearch) IsVideo(video bool) *SubmissionSearch {
	q.v.Set("is_video", strconv.FormatBool(video))
	return q
}

func (q *SubmissionSearch) Sticky(sticky bool) *SubmissionSearch {
	q.v.Set("stickied", strconv.FormatBool(sticky))
	return q
}

func (q *SubmissionSearch) CrossPostable(cross bool) *SubmissionSearch {
	q.v.Set("is_crosspostable", strconv.FormatBool(cross))
	return q
}

func (q *SubmissionSearch) CanGild(canGild bool) *SubmissionSearch {
	q.v.Set("can_gild", strconv.FormatBool(canGild))
	return q
}

func (q *SubmissionSearch) ByBot(isBot bool) *SubmissionSearch {
	q.v.Set("is_bot", strconv.FormatBool(isBot))
	return q
}

func (q *SubmissionSearch) Edited(edited bool) *SubmissionSearch {
	q.v.Set("is_edited", strconv.FormatBool(edited))
	return q
}

func (q *SubmissionSearch) RemovedByModerator(removed bool) *SubmissionSearch {
	q.v.Set("mod_removed", strconv.FormatBool(removed))
	return q
}

func (q *SubmissionSearch) RemovedByUser(removed bool) *SubmissionSearch {
	q.v.Set("user_removed", strconv.FormatBool(removed))
	return q
}

func (q *SubmissionSearch) TitleLength(i NumFilter) *SubmissionSearch {
	q.v.Set("title_length", i.String())
	return q
}

func (q *SubmissionSearch) Score(i NumFilter) *SubmissionSearch {
	q.v.Set("score", i.String())
	return q
}

func (q *SubmissionSearch) IsNoFollow(noFollow bool) *SubmissionSearch {
	q.v.Set("no_follow", strconv.FormatBool(noFollow))
	return q
}

func (q *SubmissionSearch) WantsNotifications(send bool) *SubmissionSearch {
	q.v.Set("send_replies", strconv.FormatBool(send))
	return q
}

func (q *SubmissionSearch) ID(ids ...string) *SubmissionSearch {
	addQueryList(q.v, "id", ids)
	return q
}

func (q *SubmissionSearch) Distinguished(dist bool) *SubmissionSearch {
	q.v.Set("distinguished", strconv.FormatBool(dist))
	return q
}

func (q *SubmissionSearch) SubredditType(typ SubredditType) *SubmissionSearch {
	q.v.Set("subreddit_type", typ.String())
	return q
}

func (q *SubmissionSearch) GildedSilverTimes(i NumFilter) *SubmissionSearch {
	q.v.Set("gid_1", i.String())
	return q
}

func (q *SubmissionSearch) GildedGoldTimes(i NumFilter) *SubmissionSearch {
	q.v.Set("gid_2", i.String())
	return q
}

func (q *SubmissionSearch) GildedPlatinumTimes(i NumFilter) *SubmissionSearch {
	q.v.Set("gid_3", i.String())
	return q
}

func (q *SubmissionSearch) GildedTimes(i NumFilter) *SubmissionSearch {
	q.v.Set("gilded", i.String())
	return q
}

func (q *SubmissionSearch) IsGilded(gilded bool) *SubmissionSearch {
	if gilded {
		return q.GildedTimes(GreaterThan(0))
	} else {
		return q.GildedTimes(0)
	}
}

func (q *SubmissionSearch) AuthorFlair(flair string) *SubmissionSearch {
	q.v.Set("author_flair_text", flair)
	return q
}

func (q *SubmissionSearch) LinkFlair(flair string) *SubmissionSearch {
	q.v.Set("link_flair_text", flair)
	return q
}

func (q *SubmissionSearch) AsChart() *ChartQuery {
	q.v.Set("output", "png")
	return &ChartQuery{
		v:       q.v,
		baseURL: submissionSearchBaseURL,
	}
}
