package pushshift

import (
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ChartQuery struct {
	v url.Values
	baseURL string
}

func (q *ChartQuery) ScientificNotation(sci bool) *ChartQuery {
	q.v.Set("chart.set_scientific", strconv.FormatBool(sci))
	return q
}

func (q *ChartQuery) ColorMap(c ColorMap) *ChartQuery {
	q.v.Set("chart.colormap", c.String())
	return q
}

func (q *ChartQuery) ChartTitle(str string) *ChartQuery {
	q.v.Set("chart.title", str)
	return q
}

func (q *ChartQuery) ChartSubtitle(str string) *ChartQuery {
	q.v.Set("chart.subtitle", str)
	return q
}

func (q *ChartQuery) ChartSubtitleFontSize(px uint) *ChartQuery {
	q.v.Set("chart.subtitle.fontsize", strconv.FormatUint(uint64(px), 10))
	return q
}

func (q *ChartQuery) Sorted() *ChartQuery {
	q.v.Set("chart.sort", "true")
	return q
}

func (q *ChartQuery) ChartXLogarithmic(log bool) *ChartQuery {
	q.v.Set("chart.xaxis.log", strconv.FormatBool(log))
	return q
}

func (q *ChartQuery) ChartXScale(scale float64) *ChartQuery {
	q.v.Set("chart.xaxis.scale", strconv.FormatFloat(scale, 'f', 2, 64))
	return q
}

func (q *ChartQuery) ChartYLogarithmic(log bool) *ChartQuery {
	q.v.Set("chart.yaxis.log", strconv.FormatBool(log))
	return q
}

func (q *ChartQuery) ChartBucketSize(size float64) *ChartQuery {
	q.v.Set("chart.bucket.size", strconv.FormatFloat(size, 'f', 2, 64))
	return q
}

func (q *ChartQuery) ChartYMin(min float64) *ChartQuery {
	q.v.Set("chart.ymin", strconv.FormatFloat(min, 'f', 2, 64))
	return q
}

func (q *ChartQuery) ChartYMax(min float64) *ChartQuery {
	q.v.Set("chart.ymax", strconv.FormatFloat(min, 'f', 2, 64))
	return q
}

func (q *ChartQuery) ChartXMin(min float64) *ChartQuery {
	q.v.Set("chart.xmin", strconv.FormatFloat(min, 'f', 2, 64))
	return q
}

func (q *ChartQuery) ChartXLabel(str string) *ChartQuery {
	q.v.Set("chart.xlabel", str)
	return q
}

func (q *ChartQuery) ChartYLabel(str string) *ChartQuery {
	q.v.Set("chart.ylabel", str)
	return q
}

func (q *ChartQuery) Trim() *ChartQuery {
	q.v.Set("chart.trim", "true")
	return q
}

func (q *ChartQuery) GetImage() (image.Image, error) {
	res, err := q.execute()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (q *ChartQuery) GetPNGBytes() ([]byte, error) {
	res, err := q.execute()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (q *ChartQuery) execute() (*http.Response, error) {
	u, err := url.Parse(q.baseURL)
	if err != nil {
		return nil, err
	}

	u.RawQuery = q.v.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	setHTTPHeaders(&req.Header)

	rateLimit()

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		res.Body.Close()
		return nil, fmt.Errorf("HTTP status: %s", res.Status)
	}

	if res.Header.Get("Content-Type") != "image/png" {
		res.Body.Close()
		return nil, fmt.Errorf("unexpected data: %s",
			res.Header.Get("Content-Type"))
	}

	return res, nil
}
